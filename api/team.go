package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateTeam(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create team"})
	log.Info("Start creating team")

	log.Info("Binding request")
	var request model.TeamModel
	if err := c.BindJSON(&request); err != nil {
		log.Errorf("Error during binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else if request.Save(); err != nil {
		log.Errorf("Error during saving team: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during saving team",
			Error:   err.Error(),
		})
	} else {
		log.Info("Team saved successfully")
		c.JSON(http.StatusCreated, model.CreateTeamResponse{
			Id:   request.ID,
			Name: request.Name,
		})
	}

}

func UpdateTeam(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Update team"})
	log.Info("Start updating team")

	id, isOk := getIdFromGin(c, TeamId)
	if !isOk {
		return
	}
	if team, err := model.GetTeam(id); err != nil {
		log.Errorf("Error during getting team: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting team member",
			Error:   err.Error(),
		})
	} else {
		log.Info("Binding request")
		var updateRequest model.UpdateTeamRequest
		if err := c.BindJSON(&updateRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if err := team.Update(&updateRequest); err != nil {
			log.Errorf("Error during updating team: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during updating team",
				Error:   err.Error(),
			})
		} else {
			log.Info("Team updated")
			c.Status(http.StatusOK)
		}
	}

}

func DeleteTeam(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Delete team"})
	log.Info("Start deleting team")

	id, isOk := getIdFromGin(c, TeamId)
	if !isOk {
		return
	}
	if team, err := model.GetTeam(id); err != nil {
		log.Errorf("Error during getting team: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting team",
			Error:   err.Error(),
		})
	} else {
		if err := team.Delete(); err != nil {
			log.Errorf("Error during deleting team: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during deleting team",
				Error:   err.Error(),
			})
		} else {
			log.Info("Team deleted successfully")
			c.Status(http.StatusOK)
		}
	}

}

func ListTeams(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List teams"})
	log.Info("Start listing team")

	if teams, err := model.GetTeams(); err != nil {
		log.Errorf("Error during listing teams: %s", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error during listing teams",
			Error:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, model.ConvertTeamModelToResponse(teams))
	}

}
