package db

import (
	"fmt"
	"log"

	"github.com/mnuddindev/betterkeep/auth"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", auth.Config("DB_HOST"), auth.Config("DB_USER"), auth.Config("DB_PASSWORD"), auth.Config("DB_NAME"), auth.Config("DB_PORT"))
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	utils.CheckError(err, "database connection failed")
	connection.Logger = logger.Default.LogMode(logger.Info)
	log.Println("auto migration running...")
	connection.AutoMigrate(&models.Users{})
	DB = DBInstance{
		Db: connection,
	}
}
