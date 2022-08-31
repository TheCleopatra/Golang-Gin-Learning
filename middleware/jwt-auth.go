package middleware

import (
	"github/golang-gin-learning/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// BUG-FIXED: Note here, must include versioning information
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if err != nil || token == nil || !token.Valid {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {

			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		}

	}
}
