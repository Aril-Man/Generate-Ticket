package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		privateKey := c.GetHeader("PrivateKey")
		if privateKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Private Key is required"})
			c.Abort()
			return
		}

		if privateKey != os.Getenv("PRIVATE_KEY") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Private key is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
