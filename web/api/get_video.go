package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

type VideoResponse struct{
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

func GetVideo() echo.HandlerFunc{
	return func (c echo.Context) error{
		yts := c.Get("yts").(*youtube.Service)

		videoID := c.Param("id")

		call := yts.Videos.List("id,snippet").Id(videoID)
		res, err := call.Do()
		if err != nil{
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		v := VideoResponse{VideoList:res}
		return c.JSON(http.StatusOK, v)
	}
}
