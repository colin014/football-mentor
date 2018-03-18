package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetMobileData(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Get Mobile data"})

	log.Info("Load players")
	players, err := getAllPlayer()
	if err != nil {
		log.Errorf("Error during listing players: %s", err.Error())
	}

	log.Info("Load club info")
	club, err := getClubInfo()
	if err != nil {
		log.Errorf("Error during getting club info: %s", err.Error())
	}

	resp := model.GetMobileData{
		NextGame: model.Game{
			IsHome:           true,
			OpponentTeamName: "Atlético de Madrid",
			OpponentTeamLogo: "http://en.atleticodemadrid.com/system/escudos/303/original/escudo_atm.png?1499414948",
			Date:             "20180114",
			Time:             "21:00",
		},
		Teams: []model.Team{
			{
				Name:    "Felnott",
				Players: players,
			},
		},
		Games: []model.Game{
			{
				IsHome:           true,
				OpponentTeamName: "Manchester City",
				OpponentTeamLogo: "http://pluspng.com/img-png/manchester-city-fc-png-manchester-city-supporters-club-logo-manchester-city-logo-png-410.png",
				Date:             "20180114",
				Time:             "17:00",
				Result: &model.Result{
					HomeGoal: 4,
					AwayGoal: 3,
					Events: []model.Event{
						{
							IsHome:     true,
							Type:       model.Goal,
							Minute:     9,
							PlayerName: "Alex Oxlade-Chamberlain",
						},
						{
							IsHome:     false,
							Type:       model.Goal,
							Minute:     40,
							PlayerName: "Leroy Sané",
						},
						{
							IsHome:     true,
							Type:       model.Goal,
							Minute:     59,
							PlayerName: "Roberto Firmino",
						},
						{
							IsHome:     true,
							Type:       model.YellowCard,
							Minute:     60,
							PlayerName: "Roberto Firmino",
						}, {
							IsHome:     true,
							Type:       model.Goal,
							Minute:     61,
							PlayerName: "Sadio Mané",
						}, {
							IsHome:     false,
							Type:       model.YellowCard,
							Minute:     65,
							PlayerName: "Nicolás Otamendi",
						}, {
							IsHome:     true,
							Type:       model.Goal,
							Minute:     68,
							PlayerName: "Mohamed Salah",
						},
						{
							IsHome:     false,
							Type:       model.YellowCard,
							Minute:     69,
							PlayerName: "Raheem Sterling",
						},
						{
							IsHome:     false,
							Type:       model.YellowCard,
							Minute:     72,
							PlayerName: "Fernandinho",
						},
						{
							IsHome:     false,
							Type:       model.Goal,
							Minute:     84,
							PlayerName: "Bernardo Silva",
						},
						{
							IsHome:     false,
							Type:       model.Goal,
							Minute:     90,
							PlayerName: "Ilkay Gündogan",
						},
						{
							IsHome:     true,
							Type:       model.YellowCard,
							Minute:     90,
							PlayerName: "James Milner",
						},
					},
				},
			},
			{
				IsHome:           false,
				OpponentTeamName: "Crystal Palace FC",
				OpponentTeamLogo: "https://upload.wikimedia.org/wikipedia/hif/c/c1/Crystal_Palace_FC_logo.png",
				Date:             "20180331",
				Time:             "13:30",
			},
		},
	}

	if club != nil {
		resp.TeamName = club.Name
		resp.TeamLogoUrl = club.LogoUrl
		resp.LeagueLogoUrl = club.LeagueLogoUrl
		resp.CurrentPlace = club.CurrentPlace
		resp.StadiumName = club.StadiumName
		resp.LeagueName = club.LeagueName
		resp.FacebookPageId = club.FacebookPageId
		resp.Website = club.Website
	}

	c.JSON(http.StatusOK, resp)

}
