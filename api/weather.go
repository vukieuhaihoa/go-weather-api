package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	db "github.com/vukieuhaihoa/go-weather-api/db/sqlc"
	"github.com/vukieuhaihoa/go-weather-api/services"
)

type getTemperatureRequest struct {
	NameofCity string `uri:"city" binding:"required"`
}

func (s *Server) getTemperature(ctx *gin.Context) {
	var req getTemperatureRequest
	var data services.OpenWeatherMapData

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	weatherInfor, err := s.store.RedisClient.Get(context.Background(), req.NameofCity).Result()

	if err == redis.Nil {
		// key does not exist
		data, err = services.GetWeatherByCityID(req.NameofCity)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		// store result to redis
		s.store.RedisClient.Set(context.Background(), req.NameofCity, jsonData, 30*time.Second)

	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		// Exist key on redis server

		err := json.Unmarshal([]byte(weatherInfor), &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

	}

	if data == (services.OpenWeatherMapData{}) {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("not found")))
		return
	}

	// check location stored in db ?
	location, err := s.store.GetLocationByName(context.Background(), req.NameofCity)
	if err != nil {
		// This is the first time access to city => store data to db
		if err == sql.ErrNoRows {

			arg := db.CreateLocationParams{
				Name:      req.NameofCity,
				Longitude: data.Coord.Lon,
				Latitude:  data.Coord.Lat,
			}

			_, err := s.store.CreateLocation(context.Background(), arg)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusOK, data)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// City have stored in db => update count
	_, err = s.store.UpdateLocation(context.Background(), location.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}
