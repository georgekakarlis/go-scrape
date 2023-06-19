package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goscrape.com/database"
	"goscrape.com/models"
)


func RequireAuth( router *gin.Context) {
	//get cookie of request
	tokenString, err := router.Cookie("Authorization")

	if err != nil {
		router.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//check the expiration
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		router.AbortWithStatus(http.StatusUnauthorized)
	}
	//find the user with token sub
	var user models.User
	database.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		router.AbortWithStatus(http.StatusUnauthorized)
	}
	//attach to request
	router.Set("user", user)

	//continue
	router.Next()
	fmt.Println(claims["foo"], claims["nbf"])
	} else {
	router.AbortWithStatus(http.StatusUnauthorized)
	}


}
