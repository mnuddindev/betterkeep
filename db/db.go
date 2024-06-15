package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mnuddindev/betterkeep/auth"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	var err error
	p := auth.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	utils.CheckError(err, "Error parsing str to int db")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", auth.Config("DB_USER"), auth.Config("DB_PASSWORD"), port, auth.Config("DB_NAME"))
	connection, err := gorm.Open(mysql.New(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.CheckError(err, "database connection failed")
	connection.Logger = logger.Default.LogMode(logger.Info)
	log.Println("auto migration running...")
	connection.AutoMigrate(&models.Users{})
	DB = DBInstance{
		Db: connection,
	}
}
