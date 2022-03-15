package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vukieuhaihoa/go-weather-api/db/sqlc"
)

type createLocationRequest struct {
	Name      string  `json:"name" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}

func (s *Server) createLocation(ctx *gin.Context) {
	var req createLocationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLocationParams{
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	location, err := s.store.CreateLocation(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, location)
}
