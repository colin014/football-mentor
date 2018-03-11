package api

import (
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/database"
	"github.com/colin014/football-mentor/model"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/config"
	"github.com/jinzhu/gorm"
	"strconv"
)

var logger *logrus.Logger
var db *gorm.DB

func init() {
	logger = config.Logger()
	db = database.GetDB()
}

func GetPlayers(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List players"})
	log.Info("Start getting players")

	var players []model.Player

	db := database.GetDB()
	if err := db.Find(&players).Error; err != nil {
		log.Errorf("Error during loading players from database: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during getting players from database",
			Error:   err.Error(),
		})

	} else {
		log.Info("getting players from database succeeded")
		c.JSON(http.StatusOK, players)
	}

}

func CreatePlayer(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create players"})
	log.Info("Start creating player")

	log.Info("Binding request")
	var playerRequest model.Player
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
		if err := db.Save(&playerRequest).Error; err != nil {
			log.Errorf("Error during save player: %s", err.Error())
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error during save",
				Error:   err.Error(),
			})
		} else {
			log.Info("Save succeeded")
			c.Status(http.StatusOK)
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
		var player model.Player
		if err := db.Where(model.Player{Model: gorm.Model{ID: uint(id)}}).First(&player).Error; err != nil {
			log.Errorf("Error during getting player from database[%d]: %s", uint(id), err.Error())
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Player not found",
				Error:   err.Error(),
			})
		} else {
			if err := db.Delete(player).Error; err != nil {
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
