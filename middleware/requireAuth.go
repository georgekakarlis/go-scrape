package middleware

/* import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goscrape.com/database"
	"goscrape.com/models"
)

 func RequireAuth(router *gin.Context) {


	// Print the request headers
	fmt.Println("Request Headers:")
	for key, value := range router.Request.Header {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("in middleware")
	// Get the token from the cookie
	tokenString, err := router.Cookie("Authorization")
	if err != nil {
		router.AbortWithStatus(http.StatusUnauthorized)

		return
	}

// Parse the token
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {


    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    fmt.Println("error 2")
    return []byte(os.Getenv("SECRET")), nil
})

if err != nil || !token.Valid {

    router.AbortWithStatus(http.StatusUnauthorized)
    return
}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
if !ok {

    router.AbortWithStatus(http.StatusUnauthorized)
    return
}

	// Check token expiration
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		fmt.Println("Token expired or invalid")
		router.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Retrieve user from database
	var user models.User
	database.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		router.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach user to the request context
	router.Set("user", user)

	// Continue with the next middleware or handler
	router.Next()
}


*/