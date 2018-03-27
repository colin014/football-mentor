package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/banzaicloud/banzai-types/components"
	"github.com/colin014/football-mentor/utils"
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

func CreateEvents(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Add events"})
	log.Info("Start add events")

	log.Info("Binding request")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	var createEventRequest model.CreateEventRequest
	if err := c.BindJSON(&createEventRequest); err != nil {
		log.Errorf("Error during bind json: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding",
			Error:   err.Error(),
		})
		return
	}

	log.Debugf("Binding succeeded: %#v", createEventRequest)
	log.Info("Save events into database")

	if err := createEventRequest.SaveEvents(uint(gameId)); err != nil {
		log.Errorf("Error during save events: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during save",
			Error:   err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)

}

func ListEvents(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List events"})

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	log.Info("Start listing events by gameId: %s", gameId)

	if events, err := model.GetAllEvents(uint(gameId)); err != nil {
		log.Errorf("Error during listing events: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during listing events",
			Error:   err.Error(),
		})
	} else {
		log.Info("Load events succeeded")
		c.JSON(http.StatusOK, events)
	}

}

func getGameId(c *gin.Context) (int, bool) {
	gameId, err := utils.ConvertStringToInt(c.Param("gameid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "GameId is not a number",
			Error:   "Wrong game id",
		})
		return 0, false
	} else {
		return int(gameId), true
	}
}
