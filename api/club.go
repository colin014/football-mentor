package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetClubInfo(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Get club info"})

	if club, err := getClubInfo(); err != nil {
		log.Errorf("Error during getting club info: %s", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error during getting club info",
			Error:   err.Error(),
		})
	} else {
		log.Info("Getting club info succeeded")
		c.JSON(http.StatusOK, club)
	}

}

func UpdateClubInfo(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Update club info"})

	log.Info("Binding request")
	var club model.Club
	if err := c.BindJSON(&club); err != nil {
		log.Error("Error during binding request")
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else {
		log.Debugf("Binding request succeeded: %v", club)
		log.Info("Save to database")

		var savedInfo model.Club
		db.Where(model.Club{Model: gorm.Model{ID: 1}}).First(&savedInfo)
		savedInfo.Update(&club)

		if err := db.Save(&savedInfo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error during save club info ",
				Error:   err.Error(),
			})
		} else {
			log.Info("Saving succeeded")
			c.JSON(http.StatusOK, savedInfo)
		}
	}

}

func getClubInfo() (*model.Club, error) {

	var club model.Club

	if err := db.Find(&club).Error; err != nil {
		return nil, err
	}
	return &club, nil

}
