package handlers

import (
    "NativeCloud_HR/service"
    "net/http"
    "github.com/gin-gonic/gin"
)

// 登入的請求格式
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// 登入邏輯
func LoginHandler(c *gin.Context) {
    var req LoginRequest

    // 解析 JSON 請求
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // 驗證
    token, err := service.AuthenticateUser(req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
    })
}

// 登出邏輯
func LogoutHandler(c *gin.Context) {
    // 模擬登出流程（JWT 可以由前端刪除，或加入黑名單）
    c.JSON(http.StatusOK, gin.H{
        "message": "Logout successful",
    })
}
