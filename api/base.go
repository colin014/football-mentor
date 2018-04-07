package api

import (
	"github.com/colin014/football-mentor/config"
	"github.com/colin014/football-mentor/database"
	"github.com/colin014/football-mentor/model"
	"github.com/colin014/football-mentor/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
)

var logger *logrus.Logger
var db *gorm.DB

func init() {
	logger = config.Logger()
	db = database.GetDB()
}

type GinParam string

var (
	GameId   GinParam = "gameid"
	PlayerId GinParam = "playerid"
	StaffId  GinParam = "staffid"
	EventId  GinParam = "eventid"
)

func getIdFromGin(c *gin.Context, param GinParam) (uint, bool) {
	if gameId, err := utils.ConvertStringToInt(c.Param(string(param))); err != nil {
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
