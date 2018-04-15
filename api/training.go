package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateTraining(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create training"})
	log.Info("Start creating training")

	log.Info("Binding request")
	var request model.TrainingModel
	if err := c.BindJSON(&request); err != nil {
		log.Errorf("Error during binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else if err := request.Validate(); err != nil {
		log.Errorf("Error during validate request: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during validate request",
			Error:   err.Error(),
		})
	} else if request.Save(); err != nil {
		log.Errorf("Error during saving training: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during saving training",
			Error:   err.Error(),
		})
	} else {
		log.Info("Training saved successfully")
		c.JSON(http.StatusCreated, model.CreateTrainingResponse{
			Id: request.ID,
		})
	}

}

func ListTrainings(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List trainings"})
	log.Info("Start listing trainings")

	if training, err := model.GetTrainings(); err != nil {
		log.Errorf("Error during listing trainings: %s", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error during listing training",
			Error:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, model.ConvertTrainingsModelToResponse(training))
	}

}

func UpdateTraining(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Update training"})
	log.Info("Start updating training")

	id, isOk := getIdFromGin(c, TrainingId)
	if !isOk {
		return
	}

	log.Info("Getting training from database")

	if training, err := model.GetTraining(id); err != nil {
		log.Errorf("Error during getting training: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting training",
			Error:   err.Error(),
		})
	} else {
		log.Info("Getting training succeeded")
		var updateRequest model.UpdateTrainingRequest
		if err := c.BindJSON(&updateRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if err := updateRequest.Validate(); err != nil {
			log.Errorf("Error during validate request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during validate request",
				Error:   err.Error(),
			})
		} else if err := training.Update(&updateRequest); err != nil {
			log.Errorf("Error during updating training: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during updating training",
				Error:   err.Error(),
			})
		} else {
			log.Info("Training updated")
			c.Status(http.StatusOK)
		}

	}

}

func DeleteTraining(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Delete training"})
	log.Info("Start deleting training")

	trainingId, isOK := getIdFromGin(c, TrainingId)
	if !isOK {
		return
	}

	if training, err := model.GetTraining(trainingId); err != nil {
		log.Errorf("Error during getting training: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting training",
			Error:   err.Error(),
		})
	} else if err := training.Delete(); err != nil {
		log.Errorf("Error during deleting training: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during deleting training",
			Error:   err.Error(),
		})
	} else {
		log.Info("Training deleted successfully")
		c.Status(http.StatusOK)
	}

}
