package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/banzaicloud/banzai-types/components"
	"github.com/colin014/football-mentor/utils"
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
		log.Errorf("Error during binding request: %s", err.Error())
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

func UpdatePlayer(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Update player"})
	log.Info("Start updating player")

	playerId, isOk := getPlayerId(c)
	if !isOk {
		return
	}

	if player, err := model.GetPlayer(uint(playerId)); err != nil {
		log.Errorf("Error during getting player: %s", err.Error())
		c.JSON(http.StatusNotFound, components.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Player not found",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request")
		var playerRequest model.UpdatePlayerRequest
		if err := c.BindJSON(&playerRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if err := player.Update(&playerRequest); err != nil {
			log.Errorf("Error during update player: %s", err.Error())
			c.JSON(http.StatusBadRequest, components.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during update player",
				Error:   err.Error(),
			})
		} else {
			log.Info("Player updated successfully")
			c.Status(http.StatusAccepted)
		}
	}

}

func DeletePlayer(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Delete player"})

	id, isOk := getPlayerId(c)
	if !isOk {
		return
	}

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

func getPlayerId(c *gin.Context) (int, bool) {
	gameId, err := utils.ConvertStringToInt(c.Param("playerid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, components.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "PlayerId is not a number",
			Error:   "Wrong player id",
		})
		return 0, false
	} else {
		return int(gameId), true
	}
}
