package routes

import (
	handlers "github.com/4040www/NativeCloud_HR/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(auth *gin.RouterGroup) {
	auth.OPTIONS("/login", func(c *gin.Context) {
		c.Status(204)
	})

	auth.POST("/login", handlers.LoginHandler)
	auth.POST("/logout", handlers.LogoutHandler)

}

// curl -X GET http://localhost:8080/api/auth/login \
// -H "Content-Type: application/json" \
// -d '{
// "email": "test@example.com",
// "password": "your_password"}'
