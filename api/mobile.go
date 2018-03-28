package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/colin014/football-mentor/model"
	"net/http"
)

func GetDashboardData(c *gin.Context) {

	log := logger.WithFields(logrus.Fields{"tag": "Get dashboard"})
	log.Info("Start getting dashboard")

	log.Info("Stat getting club info")
	if club, err := getClubInfo(); err != nil {
		log.Errorf("Error during getting club info: %s", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error during getting club info",
			Error:   err.Error(),
		})
	} else {
		log.Info("Getting club info succeeded")

		log.Info("Start getting next game")
		game, err := model.GetNextGame()
		if err != nil {
			log.Errorf("Error during getting next game: %s", err.Error())
		}

		dashboard := model.DashboardResponse{
			CurrentPlace:  club.CurrentPlace,
			ClubName:      club.Name,
			ClubLogoUrl:   club.LogoUrl,
			LeagueName:    club.LeagueName,
			LeagueLogoUrl: club.LeagueLogoUrl,
			StadiumName:   club.StadiumName,
			WebUrl:        club.Website,
			NextGame:      game,
		}

		c.JSON(http.StatusOK, dashboard)

	}

}
