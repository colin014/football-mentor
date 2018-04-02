package api

import (
	"github.com/colin014/football-mentor/config"
	"github.com/colin014/football-mentor/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/colin014/football-mentor/utils"
	"net/http"
	"github.com/colin014/football-mentor/model"
)

var logger *logrus.Logger
var db *gorm.DB

func init() {
	logger = config.Logger()
	db = database.GetDB()
}

func getIdFromGin(c *gin.Context) (uint, bool) {
	if gameId, err := utils.ConvertStringToInt(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Id is not a number",
			Error:   err.Error(),
		})
		return 0, false
	} else {
		return uint(gameId), true
	}
}
