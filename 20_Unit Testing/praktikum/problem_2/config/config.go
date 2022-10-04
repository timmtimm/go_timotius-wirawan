package config

import (
	"fmt"
	"praktikum_section_20/models"
	"praktikum_section_20/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: utils.GetConfig("DB_Username"),
		DB_Password: utils.GetConfig("DB_Password"),
		DB_Port:     utils.GetConfig("DB_Port"),
		DB_Host:     utils.GetConfig("DB_Host"),
		DB_Name:     utils.GetConfig("DB_Name"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
