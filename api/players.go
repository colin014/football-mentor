package api

import (
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/database"
	"github.com/colin014/football-mentor/model"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/config"
	"github.com/jinzhu/gorm"
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
		db.Save(&playerRequest)
	}
}
