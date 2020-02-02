package main

import (
	"github.com/IkezawaYuki/nugo-go/databases"
	"github.com/IkezawaYuki/nugo-go/models"
	"github.com/sirupsen/logrus"
)

func main(){
	db, err := databases.Connect()
	defer db.Close()

	if err != nil{
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
