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
		c.JSON(http.StatusOK, model.ConvertGameModelToResponse(games))
	}

}

func UpdateGame(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Update game"})
	log.Info("Start updating game")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	if game, err := model.GetGame(uint(gameId)); err != nil {
		log.Errorf("Error during getting game: %s", err.Error())
		c.JSON(http.StatusNotFound, components.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Game not found",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request")
		var gameRequest model.UpdateGameRequest
		if err := c.BindJSON(&gameRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if err := game.Update(&gameRequest); err != nil {
			log.Errorf("Error during update game: %s", err.Error())
			c.JSON(http.StatusBadRequest, components.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during update player",
				Error:   err.Error(),
			})
		} else {
			log.Info("Game updated successfully")
			c.Status(http.StatusAccepted)
		}
	}
}

func DeleteGame(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Delete game"})

	log.Info("Start deleting game")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	if game, err := model.GetGame(uint(gameId)); err != nil {
		log.Errorf("Error during getting game: %s", err.Error())
		c.JSON(http.StatusNotFound, components.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Game not found",
			Error:   err.Error(),
		})
	} else if err := game.DeleteGame(); err != nil {
		log.Errorf("Error during delete game: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during deleting",
			Error:   err.Error(),
		})
	} else {
		log.Info("Delete succeeded")
		c.Status(http.StatusOK)
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

	log.Info("Event(s) created successfully")
	c.Status(http.StatusCreated)

}

func ListEvents(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List events"})

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	log.Infof("Start listing events by gameId: %d", gameId)

	if events, err := model.GetAllEvents(uint(gameId)); err != nil {
		log.Errorf("Error during listing events: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during listing events",
			Error:   err.Error(),
		})
	} else {
		log.Info("Load events succeeded")
		c.JSON(http.StatusOK, model.ConvertEventModelToResponse(events))
	}

}

func DeleteEvent(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Delete events"})

	log.Info("Start deleting events")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	eventId, isOk := getEventId(c)
	if !isOk {
		return
	}

	log.Infof("Start deleting event with gameId[%d], and evenId[%d]", gameId, eventId)

	if err := model.DeleteEvent(uint(gameId), uint(eventId)); err != nil {
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during deleting",
			Error:   err.Error(),
		})
		return
	}

	log.Info("Event deleted successfully")
	c.Status(http.StatusOK)

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

func getEventId(c *gin.Context) (int, bool) {
	eventId, err := utils.ConvertStringToInt(c.Param("eventid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "EventId is not a number",
			Error:   "Wrong event id",
		})
		return 0, false
	} else {
		return int(eventId), true
	}
}

func CreateResult(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Create result"})

	log.Info("Start creating result")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	log.Debugf("Game id: %d", gameId)
	log.Info("Binding request")

	var createResultRequest model.ResultModel
	if err := c.BindJSON(&createResultRequest); err != nil {
		log.Errorf("Error during binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
		return
	}

	log.Info("Binding succeeded")
	log.Info("Save into database")

	if err := createResultRequest.SaveResult(uint(gameId)); err != nil {
		log.Errorf("Error during save: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during save",
			Error:   err.Error(),
		})
		return
	}

	log.Info("Save succeeded")
	c.Status(http.StatusCreated)

}

func UpdateResult(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Update result"})
	log.Info("Start updating result")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	if _, err := model.GetGame(uint(gameId)); err != nil {
		log.Errorf("Error during getting game: %s", err.Error())
		c.JSON(http.StatusNotFound, components.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting game",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request")
		var resultRequest model.UpdateResultRequest
		if err := c.BindJSON(&resultRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, components.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if result, err := model.GetResult(uint(gameId)); err != nil {
			log.Errorf("Error during getting result: %s", err.Error())
			c.JSON(http.StatusNotFound, components.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Error during getting result",
				Error:   err.Error(),
			})
		} else if err := result.Update(&resultRequest); err != nil {
			log.Errorf("Error during updating result: %s", err.Error())
			c.JSON(http.StatusBadRequest, components.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during updating result",
				Error:   err.Error(),
			})
		} else {
			log.Info("Result updated successfully")
			c.Status(http.StatusAccepted)
		}
	}
}

func DeleteResult(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Delete result"})

	log.Info("Start deleting result")

	gameId, isOk := getGameId(c)
	if !isOk {
		return
	}

	log.Debugf("Game id: %d", gameId)

	if err := model.DeleteResult(uint(gameId)); err != nil {
		log.Errorf("Error deleting result: %s", err.Error())
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during deleting",
			Error:   err.Error(),
		})
		return
	}

	log.Info("Delete succeeded")
	c.Status(http.StatusOK)

}
