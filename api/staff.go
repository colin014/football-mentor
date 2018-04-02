package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/model"
	"net/http"
	"github.com/colin014/football-mentor/utils"
)

func ListStaff(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "List staff"})
	log.Info("Start listing staff")

	if staff, err := model.GetStaff(); err != nil {
		log.Errorf("Error during listing staff: %s", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error during listing staff",
			Error:   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, model.ConvertStaffModelToResponse(staff))
	}

}

func CreateStaffMember(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Create staff member"})
	log.Info("Start creating staff member")

	log.Info("Binding request")
	var request model.StaffModel
	if err := c.BindJSON(&request); err != nil {
		log.Errorf("Error during binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during binding request",
			Error:   err.Error(),
		})
	} else if request.Save(); err != nil {
		log.Errorf("Error during saving staff member: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during saving staff member",
			Error:   err.Error(),
		})
	} else {
		log.Info("Staff saved successfully")
		c.JSON(http.StatusCreated, model.CreateStaffResponse{
			Id:   request.ID,
			Name: request.Name,
		})
	}

}

func UpdateStaffMember(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Update staff member"})
	log.Info("Start updating staff member")

	id, isOk := getStaffId(c)
	if !isOk {
		return
	}

	log.Info("Getting staff member from database")

	if staffMember, err := model.GetStaffMember(uint(id)); err != nil {
		log.Errorf("Error during getting staff member: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting staff member",
			Error:   err.Error(),
		})
	} else {
		log.Info("Getting staff member succeeded")
		var updateRequest model.UpdateStaffRequest
		if err := c.BindJSON(&updateRequest); err != nil {
			log.Errorf("Error during binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during binding request",
				Error:   err.Error(),
			})
		} else if err := staffMember.Update(&updateRequest); err != nil {
			log.Errorf("Error during updating staff member: %s", err.Error())
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Error during updating staff member",
				Error:   err.Error(),
			})
		} else {
			log.Info("Staff member updated")
			c.Status(http.StatusOK)
		}

	}

}

func DeleteStaffMember(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Delete staff member"})
	log.Info("Start deleting staff member")

	staffId, isOK := getStaffId(c)
	if !isOK {
		return
	}

	if staffMember, err := model.GetStaffMember(uint(staffId)); err != nil {
		log.Errorf("Error during getting staff member: %s", err.Error())
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Error during getting staff member",
			Error:   err.Error(),
		})
	} else if err := staffMember.Delete(); err != nil {
		log.Errorf("Error during deleting staff member: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error during deleting staff member",
			Error:   err.Error(),
		})
	} else {
		log.Info("Staff member deleted successfully")
		c.Status(http.StatusOK)
	}

}

func getStaffId(c *gin.Context) (int, bool) {
	staffId, err := utils.ConvertStringToInt(c.Param("staffid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "StaffId is not a number",
			Error:   "Wrong staff id",
		})
		return 0, false
	} else {
		return int(staffId), true
	}
}
