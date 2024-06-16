package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
	"gorm.io/gorm"
)

func RegistrationHelper(user models.Users) (models.Users, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.Users{}, err
	}
	user.Password = string(hashedPassword)
	err = DB.Db.Debug().Model(&models.Users{}).Create(&user).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return models.Users{}, errors.New("user already exists")
	}
	if err != nil {
		return models.Users{}, nil
	}
	return user, nil
}

func GetOTP(uid uuid.UUID) int64 {
	code, err := utils.GenerateOTP()
	if err != nil {
		return 0
	}
	err = DB.Db.Debug().Model(&models.Users{}).Where("ID = ?", uid).Update("verification", code).Error
	if err != nil {
		return 0
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}
	return code
}
