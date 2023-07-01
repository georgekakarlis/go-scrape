package handlers

/*
import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"goscrape.com/database"
	"goscrape.com/models"
)

func Signup(router *gin.Context) {
	//get the email/pass of req body
	var body struct {
		FirstName string
		LastName string
		Email string
		Password string
	}

	if router.Bind(&body) != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//hash password
	//If the request body is successfully bound to the body struct, the code continues to hash the password using the bcrypt package.
	// It calls bcrypt.GenerateFromPassword to generate a secure hash of the password provided in the request body.
	//The hash cost factor is set to 10, which determines the computational cost of the hash generation.
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to hash password",
		})
		return
	}


	//create the user
	user := models.User{Email: body.Email, Password: string(hash), FirstName: body.FirstName, LastName: body.LastName}
	result := database.DB.Create(&user)

	if result.Error != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to create user",
		})

		return
	}

	//respond
	router.JSON(http.StatusOK, gin.H{})
}

 func Login( router *gin.Context) {
	//get email and pass of req body
	var body struct {
		Email string
		Password string
	}

	if router.Bind(&body) != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//look up requested user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	fmt.Printf("Retrieved User: %+v\n", user)

	if user.ID == 0 {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid password",
		})
		return
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to create token",
		})
		return
	}

	//send it back
	// Set the token as an HTTP-only and secure cookie
	router.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, false)
	router.JSON(http.StatusOK, gin.H{
		"message": "hey!",
	})
}

func Logout(router *gin.Context) {
	router.SetCookie("Authorization", "", -1, "", "", false, true)
	router.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}


func Validate( router *gin.Context) {

	user, _ := router.Get("user")
	router.JSON(http.StatusOK, gin.H{
		"message": user,
	})
} */