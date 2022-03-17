package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vukieuhaihoa/go-weather-api/db/sqlc"
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

	// router for temperature
	router.GET("/temperature/:city", server.getTemperature)

	// router for location
	router.POST("/location/", server.createLocation)
	router.GET("/location/:id", server.getLocation)
	router.PUT("location/:id", server.updateLocation)

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
