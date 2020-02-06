package api

import (
	"firebase.google.com/go/auth"
	"github.com/IkezawaYuki/nugo-go/middlewares"
	"github.com/IkezawaYuki/nugo-go/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchFavoriteVideos() echo.HandlerFunc{
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)
		var user models.User
		dbs.DB.Table("users").Where(models.User{UID:token.UID}).First(&user)

		var favorites []models.Favorite
		dbs.DB.Model(&user).Related(&favorites)

		var videosIds string
		for _, f := range favorites{
			if len(videosIds) == 0{
				videosIds += f.VideoID
			}else{
				videosIds += "," + f.VideoID
			}
		}
		yts := c.Get("yts").(*youtube.Service)
		call := yts.Videos.List("id,snippet").Id(videosIds).MaxResults(10)
		res, err := call.Do()
		if err != nil{
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
