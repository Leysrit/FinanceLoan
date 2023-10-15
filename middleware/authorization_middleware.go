package middleware

import (
	"Finance/auth"
	"Finance/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractJwtFromHeader(c.Request)
		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Token tidak ada.", Error: utility.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		claims, err := auth.GetClaimsFromJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Token tidak valid.", Error: utility.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		var username = claims.Username
		c.Set("username", username)
		c.Next()
	}
}
