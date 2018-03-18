package api

import (
	"github.com/colin014/football-mentor/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMobileData(c *gin.Context) {
	// log := logger.WithFields(logrus.Fields{"tag": "Get Mobile data"})

	players, _ := getAllPlayer()

	c.JSON(http.StatusOK, model.GetMobileData{
		TeamName:       "Liverpool FC",
		TeamLogoUrl:    "http://assets3.lfcimages.com/uploads/placeholders/6683__1925__logo-125-splash-new-padded.png",
		LeagueLogoUrl:  "https://is2-ssl.mzstatic.com/image/thumb/Purple62/v4/2e/4d/eb/2e4debf9-7cbc-b796-b26c-a119e470bde6/AppIcon-1x_U007emarketing-85-220-0-5.png/246x0w.jpg",
		CurrentPlace:   2,
		StadiumName:    "Anfield Road",
		LeagueName:     "Premier League",
		FacebookPageId: "LiverpoolFC",
		Website:        "www.liverpoolfc.com",
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
				//Players: []model.Player{
				//	{
				//		Name:         "Bodgan Adam",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p45175.png",
				//		BirthDate:    "19870927",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 34,
				//	},
				//	{
				//		Name:         "Dejan Lovren",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p38454.png",
				//		BirthDate:    "19890705",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 6,
				//	},
				//	{
				//		Name:         "James Milner",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p15157.png",
				//		BirthDate:    "19860104",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 7,
				//	},
				//	{
				//		Name:         "Trent Alexander-Arnold",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p169187.png",
				//		BirthDate:    "19981007",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 66,
				//	},
				//	{
				//		Name:         "Mohamed Salah",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p118748.png",
				//		BirthDate:    "19920615",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 11,
				//	},
				//	{
				//		Name:         "Roberto Firmino",
				//		ImageUrl:     "https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/250x250/p92217.png",
				//		BirthDate:    "19911002",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 9,
				//	},
				//	{
				//		Name:         "John Doe",
				//		ImageUrl:     "",
				//		BirthDate:    "",
				//		BirthPlace:   "",
				//		Description:  "",
				//		JerseyNumber: 19,
				//	},
				//},
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
				Result:           nil,
			},
		},
	})

}
