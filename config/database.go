package config

import (
	"bwastartup/campaign"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB{
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect %s database", driver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database", driver)
	}

	db.Debug().AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{}, &transaction.Transaction{})


	return db
}