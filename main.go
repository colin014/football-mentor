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

	logger.Info("Create table(s): ",
		model.PlayerModel.TableName(model.PlayerModel{}),
		", ",
		model.Club.TableName(model.Club{}),
		", ",
		model.GameModel.TableName(model.GameModel{}),
		", ",
		model.ResultModel.TableName(model.ResultModel{}),
		", ",
		model.Event.TableName(model.Event{}),
	)

	db.AutoMigrate(
		&model.PlayerModel{},
		&model.Club{},
		&model.GameModel{},
		&model.ResultModel{},
		&model.Event{},
	)
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api")
	v1.GET("/players", api.GetPlayers)
	v1.POST("/players", api.CreatePlayer)
	v1.DELETE("/player/:id", api.DeletePlayer)
	v1.GET("/mobile/data", api.GetMobileData)
	v1.GET("/club", api.GetClubInfo)
	v1.PUT("/club", api.UpdateClubInfo)
	v1.POST("/games", api.CreateGame)
	v1.GET("/games", api.GetGames)
	v1.POST("/games/:gameid/events", api.CreateEvents)
	v1.GET("/games/:gameid/events", api.ListEvents)
	v1.DELETE("/games/:gameid/events/:eventid", api.DeleteEvent)

	router.Run(":6060")
}
