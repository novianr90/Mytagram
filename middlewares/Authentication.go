package middlewares

import (
	"final-project-hacktiv8/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_status":  "Unauthenticated",
				"error_message": err.Error(),
			})
			return
		}

		userData := verifyToken.(jwt.MapClaims)

		c.Set("userData", userData)

		c.Next()
	}
}
