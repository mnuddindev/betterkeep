package db

import (
	"fmt"
	"log"

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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", utils.Config("DB_HOST"), utils.Config("DB_USER"), utils.Config("DB_PASSWORD"), utils.Config("DB_NAME"), utils.Config("DB_PORT"))
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
