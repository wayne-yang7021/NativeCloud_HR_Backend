package service

import (
    "errors"
    "NativeCloud_HR/model"
    "NativeCloud_HR/repository"
    "NativeCloud_HR/utils"
)

// 用 email 和密碼驗證登入，成功回傳使用者資訊與 JWT token
func AuthenticateUser(email, password string) (*model.User, string, error) {
    // 查找使用者
    user, err := repository.FindUserByEmail(email)
    if err != nil || user == nil {
        return nil, "", errors.New("User not found")
    }

    // 檢查密碼
    if !utils.CheckPasswordHash(password, user.PasswordHash) {
        return nil, "", errors.New("Invalid credentials")
    }

    // 產生 JWT token
    token, err := utils.GenerateJWT(user)
    if err != nil {
        return nil, "", errors.New("Failed to generate token")
    }

    return user, token, nil
}
