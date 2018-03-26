package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetPlayers(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List players"})
	log.Info("Start getting players")

	if players, err := model.GetAllPlayer(); err != nil {
		log.Errorf("Error during loading players from database: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during getting players from database",
			Error:   err.Error(),
		})

	} else {
		log.Info("getting players from database succeeded")
		c.JSON(http.StatusOK, model.ConvertPlayerModelToResponse(players))
	}

}

func CreatePlayer(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create players"})
	log.Info("Start creating player")

	log.Info("Binding request")
	var playerRequest model.PlayerModel
	if err := c.BindJSON(&playerRequest); err != nil {
		log.Error("Error during binding request")
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request succeeded")
		log.Info("Save player to database")
		if err := playerRequest.Save(); err != nil {
			log.Errorf("Error during save player: %s", err.Error())
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error during save",
				Error:   err.Error(),
			})
		} else {
			log.Info("Save succeeded")
			c.JSON(http.StatusCreated, model.CreatePlayerResponse{
				Id:   playerRequest.ID,
				Name: playerRequest.Name,
			})
		}
	}
}

func DeletePlayer(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Delete player"})

	playerId := c.Param("id")
	if id, err := strconv.ParseUint(playerId, 10, 32); err != nil {
		log.Errorf("Error during parsing id: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during parsing id",
			Error:   err.Error(),
		})
	} else {
		log.Infof("Player id: %d", uint(id))
		var player model.PlayerModel
		if err := model.GetPlayerById(uint(id), &player); err != nil {
			log.Errorf("Error during getting player from database[%d]: %s", uint(id), err.Error())
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Player not found",
				Error:   err.Error(),
			})
		} else {
			if err := player.Delete(); err != nil {
				log.Errorf("Error during deleting player[%d]: %s", uint(id), err.Error())
				c.JSON(http.StatusBadRequest, model.ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: "Error during deleting player",
					Error:   err.Error(),
				})
			} else {
				c.Status(http.StatusOK)
			}
		}

	}

}
