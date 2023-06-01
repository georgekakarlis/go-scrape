package handlers

import (
	"errors"
	"log"
	"net/mail"
	"time"

	"goscrape.com/config"
	"goscrape.com/database"
	"goscrape.com/models"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//This function compares a password string with its corresponding hash value.
//It uses bcrypt's CompareHashAndPassword function to compare the provided password and hash.
//If the comparison is successful and there is no error, it returns true. Additionally, it logs the provided hash value.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	log.Println(hash, "haaaash")
	return err == nil
}
//This function retrieves a user from the database based on their email address (e). 
//It queries the database to find a user with the matching email. 
//If a user is found, it returns a pointer to the models.User object representing the user. 
//If no user is found, it returns nil, nil. If any other error occurs during the retrieval process, it returns the error.
func getUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

//This function retrieves a user from the database based on their username (u). 
//It queries the database to find a user with the matching username. If a user is found, it returns a pointer to the models.User object representing the user. 
//If no user is found, it returns nil, nil. If any other error occurs during the retrieval process, it returns the error.
func getUserByUsername(u string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

//This function checks whether the provided email address (email) is valid. It uses the mail.ParseAddress function to parse the email address. 
//If parsing is successful and there is no error, it returns true, indicating that the email is valid.
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

//This function is an HTTP request handler for user login. It expects a request body with the fields identity (which can be either an email or a username) and password. 
//It parses the request body to retrieve the login input. It then determines whether the identity represents an email or a username by validating it using the valid function. 
//Based on the type of identity, it calls either getUserByEmail or getUserByUsername to retrieve the user from the database. 
//If a user is found, it compares the provided password with the user's hashed password using the CheckPasswordHash function. 
//If the password is valid, it generates a JWT token, sets the token claims (including the username, user ID, and expiration time), signs the token using the secret from the configuration, and returns the token in a JSON response. 
//If any errors occur during the login process, appropriate error responses are returned.
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}

	identity := input.Identity
	pass := input.Password
	user, email, err := new(models.User), new(models.User), *new(error)

	if valid(identity) {
		email, err = getUserByEmail(identity)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "errors": err.Error()})
		}
		ud = UserData{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	} else {
		user, err = getUserByUsername(identity)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "errors": err.Error()})
		}
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	if email == nil && user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "errors": err.Error()})
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}