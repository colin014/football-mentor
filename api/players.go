package api

import (
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/database"
	"github.com/colin014/football-mentor/model"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/config"
)

var logger *logrus.Logger

func init() {
	logger = config.Logger()
}

func GetPlayers(c *gin.Context) {

	logger.Info("Start getting players")

	var players []model.Player

	db := database.GetDB()
	if err := db.Find(&players).Error; err != nil {
		logger.Errorf("Error during loading players from database: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during getting players from database",
			Error:   err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, players)
	}

}
