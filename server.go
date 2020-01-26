package main

import (
	"github.com/IkezawaYuki/nugo-go/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init(){
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
