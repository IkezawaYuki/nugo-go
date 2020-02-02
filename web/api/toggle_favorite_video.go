package api

import "github.com/labstack/echo"

type ToggleFavoriteVideoResponse struct{
	VideoID string `json:"video_id"`
	IsFavorite bool `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc{
	return func(c echo.Context) error {
		// todo 
	}
}