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
	var gameRequest model.Game
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
			log.Info("Game save succeeded")
			log.Info("Save event(s)")

			if gameRequest.Result != nil {
				for _, event := range gameRequest.Result.Events {
					event.ResultId = gameRequest.Result.GameId
					if err := db.Save(&event).Error; err != nil {
						log.Errorf("Error during save event [%v]: %s", event, err.Error())
						c.JSON(http.StatusInternalServerError, model.ErrorResponse{
							Code:    http.StatusInternalServerError,
							Message: "Error during save event(s)",
							Error:   err.Error(),
						})
						return
					}
				}
			}
			c.Status(http.StatusCreated)
		}
	}

}

func GetGames(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List games"})
	log.Info("Start getting games")

	if games, err := getAllGames(); err != nil {
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

func getAllGames() ([]model.Game, error) {
	log := logger.WithFields(logrus.Fields{"tag": "Load games from database"})
	var games []model.Game

	if err := db.Find(&games).Error; err != nil {
		return nil, err
	}

	for i, game := range games {
		games[i].Result = &model.Result{}
		if err := db.Where(model.Result{GameId: game.ID}).First(&(*games[i].Result)).Error; err != nil {
			log.Warnf("Error during load result: %s", err.Error())
			games[i].Result = nil
		} else {
			var events []model.Event
			if err := db.Where(model.Event{ResultId: games[i].Result.GameId}).Find(&events).Error; err != nil {
				log.Warnf("Error during load events to events: %s", err.Error())
				games[i].Result.Events = nil
			} else if len(events) != 0 {
				games[i].Result.Events = events
			}
		}
	}

	return games, nil
}
