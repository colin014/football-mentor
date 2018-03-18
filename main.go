package main

import (
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/api"
	"github.com/colin014/football-mentor/database"
	"github.com/colin014/football-mentor/model"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/config"
)

var logger *logrus.Logger

func init() {
	logger = config.Logger()

	db := database.GetDB()

	logger.Infof("Create %s table(s)", model.Player.TableName(model.Player{}))

	db.AutoMigrate(&model.Player{})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api")
	v1.GET("/players", api.GetPlayers)
	v1.POST("/players", api.CreatePlayer)
	v1.DELETE("/player/:id", api.DeletePlayer)
	v1.GET("/mobile/data", api.GetMobileData)
	router.Run(":6060")
}
