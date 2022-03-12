package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-weather-api/services"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("temperature/:city", func(ctx *gin.Context) {
		city := ctx.Param("city")
		fmt.Println(city)
		data, err := services.GetWeatherByCityID(city)

		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, data)
	})
}
