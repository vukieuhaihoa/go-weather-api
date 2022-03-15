package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vukieuhaihoa/go-weather-api/db/sqlc"
	"github.com/vukieuhaihoa/go-weather-api/services"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// router for temperature
	router.GET("/temperature/:city", func(ctx *gin.Context) {
		city := ctx.Param("city")

		data, err := services.GetWeatherByCityID(city)

		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusOK, data)
	})

	// router for location
	router.POST("/location", server.createLocation)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
