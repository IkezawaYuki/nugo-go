package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect()(db *gorm.DB,err error){
	err = godotenv.Load()
	if err != nil{
		logrus.Fatal("Error loading .env file")
	}

	// todo 

	return db, err
}
