package routes

import (
    "github.com/gin-gonic/gin"
    "NativeCloud_HR/api/handlers"
)

func RegisterRoutes(auth *gin.RouterGroup) {

    auth.POST("/login", handlers.LoginHandler)
    auth.POST("/logout", handlers.LogoutHandler)

}
