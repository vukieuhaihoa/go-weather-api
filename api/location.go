package api

import (
	"context"
	"database/sql"
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

type getLocationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getLocation(ctx *gin.Context) {
	var req getLocationRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	location, err := s.store.GetLocation(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, location)
}

func (s *Server) updateLocation(ctx *gin.Context) {
	var req getLocationRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	location, err := s.store.UpdateLocation(context.Background(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, location)
}
