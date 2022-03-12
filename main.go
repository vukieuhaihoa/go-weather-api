package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-weather-api/handlers"
)

func main() {
	r := gin.Default()
	handlers.InitRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
