package api

import (
	"firebase.google.com/go/auth"
	"github.com/IkezawaYuki/nugo-go/middlewares"
	"github.com/IkezawaYuki/nugo-go/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

type VideoResponse struct{
	VideoList *youtube.VideoListResponse `json:"video_list"`
	IsFavorite bool `json:"is_favorite"`
}

func GetVideo() echo.HandlerFunc{
	return func (c echo.Context) error{
		yts := c.Get("yts").(*youtube.Service)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		videoID := c.Param("id")
		isFavorite := false
		if token != nil{
			favorite := models.Favorite{}
			isFavoriteNotFound := dbs.DB.Table("favorite").
				Joins("INNER JOIN users ON users.id = favorites.user_id").
				Where(models.User{UID: token.UID}).
				Where(models.Favorite{VideoID: videoID}).
				First(&favorite).
				RecordNotFound()
			logrus.Debug("isFavoriteNotFount: ", isFavoriteNotFound)
			if !isFavoriteNotFound{
				isFavorite = true
			}
		}

		call := yts.Videos.List("id,snippet").Id(videoID)
		res, err := call.Do()
		if err != nil{
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		v := VideoResponse{
			VideoList:res,
			IsFavorite: isFavorite,
		}
		return c.JSON(http.StatusOK, v)
	}
}
