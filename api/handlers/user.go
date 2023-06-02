package handlers

import (
	"strconv"

	"goscrape.com/database"
	"goscrape.com/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//This function takes a password string as input, hashes it using bcrypt with a cost factor of 14, and returns the hashed password as a string along with any error encountered during the hashing process.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
//here it validates the provided JWT token (t) and the user ID (id). It first converts the id string to an integer and checks for any conversion errors. 
//Then, it extracts the user ID from the token claims and compares it to the converted id. 
//It returns true if the user ID from the token matches the converted id, indicating a valid token and user ID, and false otherwise.
func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}
//This function is an HTTP request handler for retrieving a user's information.
// It extracts the user ID from the request parameters (id). It queries the database to find the user with the specified ID. 
//If the user is found, it returns a JSON response with the user's information. If the user is not found, it returns a JSON response with a 404 status code and an appropriate error message.
func validUser(id string, p string) bool {
	db := database.DB
	var user models.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// Here this function is an HTTP request handler for retrieving a user's information. It extracts the user ID from the request parameters (id). 
//It queries the database to find the user with the specified ID. If the user is found, it returns a JSON response with the user's information. 
//If the user is not found, it returns a JSON response with a 404 status code and an appropriate error message.
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user models.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

//This function is an HTTP request handler for creating a new user. It parses the request body to retrieve the user information, including the username and email. 
//It performs validation on the user struct using the validator package. 
//If the validation passes, it hashes the user's password, creates a new user in the database, and returns a JSON response with the created user's information. 
//If any errors occur during the process, appropriate error responses are returned.
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	newUser := NewUser{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// Here we have an HTTP request handler for updating a user's information. It parses the request body to retrieve the updated user information, specifically the names field. 
//It also retrieves the user ID from the request parameters and the JWT token from the user context. It validates the token and user ID using the validToken function. 
//If the token and user ID are valid, it fetches the user from the database, updates the names field, saves the changes, and returns a JSON response with the updated user's information.
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
	var user models.User

	db.First(&user, id)
	user.Names = uui.Names
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

//This function is an HTTP request handler for deleting a user. It parses the request body to retrieve the user's password. 
//It also retrieves the user ID from the request parameters and the JWT token from the user context. 
//It validates the token and user ID using the validToken function and checks if the user is valid using the validUser function. If the token, user ID, and user are all valid, it deletes the user from the database and returns a JSON response indicating the success of the deletion.
func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

	}

	if !validUser(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})

	}

	db := database.DB
	var user models.User

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}