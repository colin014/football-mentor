package main

import (
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/api"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api")
	v1.GET("/players", api.GetPlayers)
	router.Run(":6060")
}
