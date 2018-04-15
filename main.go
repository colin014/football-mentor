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
		model.EventModel.TableName(model.EventModel{}),
		", ",
		model.StaffModel.TableName(model.StaffModel{}),
		", ",
		model.TeamModel.TableName(model.TeamModel{}),
		", ",
		model.TrainingModel.TableName(model.TrainingModel{}),
	)

	db.AutoMigrate(
		&model.PlayerModel{},
		&model.Club{},
		&model.GameModel{},
		&model.ResultModel{},
		&model.EventModel{},
		&model.StaffModel{},
		&model.TeamModel{},
		&model.TrainingModel{},
	)
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api")

	v1.GET("/players", api.GetPlayers)
	v1.POST("/players", api.CreatePlayer)
	v1.PUT("/players/:playerid", api.UpdatePlayer)
	v1.DELETE("/player/:playerid", api.DeletePlayer)

	v1.GET("/club", api.GetClubInfo)
	v1.PUT("/club", api.UpdateClubInfo)

	v1.POST("/games", api.CreateGame)
	v1.GET("/games", api.GetGames)
	v1.PUT("/games/:gameid", api.UpdateGame)
	v1.DELETE("/games/:gameid", api.DeleteGame)

	v1.POST("/games/:gameid/result", api.CreateResult)
	v1.PUT("/games/:gameid/result", api.UpdateResult)
	v1.DELETE("/games/:gameid/result", api.DeleteResult)

	v1.POST("/games/:gameid/events", api.CreateEvents)
	v1.GET("/games/:gameid/events", api.ListEvents)
	v1.PUT("/games/:gameid/events/:eventid", api.UpdateEvent)
	v1.DELETE("/games/:gameid/events/:eventid", api.DeleteEvent)

	v1.GET("/staff", api.ListStaff)
	v1.POST("/staff", api.CreateStaffMember)
	v1.PUT("/staff/:staffid", api.UpdateStaffMember)
	v1.DELETE("/staff/:staffid", api.DeleteStaffMember)

	v1.GET("/team", api.ListTeams)
	v1.POST("/team", api.CreateTeam)
	v1.PUT("/team/:teamid", api.UpdateTeam)
	v1.DELETE("/team/:teamid", api.DeleteTeam)

	v1.GET("training", api.ListTrainings)
	v1.POST("training", api.CreateTraining)
	v1.PUT("training/:trainingid", api.UpdateTraining)
	v1.DELETE("training/:trainingid", api.DeleteTraining)

	v1.GET("/mobile/config", api.GetConfig)
	v1.GET("/mobile/dashboard", api.GetDashboardData)

	router.Run(":6060")
}
