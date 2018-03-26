package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateGame(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create games"})
	log.Info("Start creating game")

	log.Info("Binding request")
	var gameRequest model.GameModel
	if err := c.BindJSON(&gameRequest); err != nil {
		log.Error("Error during binding request")
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request succeeded")
		log.Info("Save game to database")
		if err := db.Save(&gameRequest).Error; err != nil {
			log.Errorf("Error during save game: %s", err.Error())
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error during save",
				Error:   err.Error(),
			})
		} else {
			log.Info("GameModel save succeeded")

			c.JSON(http.StatusCreated, model.CreateGameResponse{
				Id:           gameRequest.ID,
				OpponentTeam: gameRequest.OpponentTeamName,
			})
		}
	}

}

func GetGames(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List games"})
	log.Info("Start getting games")

	if games, err := model.GetAllGames(); err != nil {
		log.Errorf("Error during loading games from database: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during getting games from database",
			Error:   err.Error(),
		})

	} else {
		log.Info("getting games from database succeeded")
		c.JSON(http.StatusOK, games)
	}

}
