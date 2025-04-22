package repository

import (
    "NativeCloud_HR/model"
    // 假設你用 GORM
    "gorm.io/gorm"
)

func FindUserByEmail(email string) (*model.User, error) {
    var user model.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
