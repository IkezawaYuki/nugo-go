package routes

import (
	"github.com/IkezawaYuki/nugo-go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo){
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}