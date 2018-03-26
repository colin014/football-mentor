package model

type GetMobileData struct {
	TeamName       string      `json:"team_name"`
	TeamLogoUrl    string      `json:"team_logo_url"`
	LeagueLogoUrl  string      `json:"league_logo_url"`
	CurrentPlace   int         `json:"current_place"`
	StadiumName    string      `json:"stadium_name"`
	LeagueName     string      `json:"league_name"`
	FacebookPageId string      `json:"facebook_page_id"`
	Website        string      `json:"website"`
	NextGame       GameModel   `json:"next_game"`
	Teams          []Team      `json:"teams"`
	Games          []GameModel `json:"games"`
}

type Team struct {
	Name    string        `json:"name"`
	Players []PlayerModel `json:"players"`
}

type EventType int

const (
	Goal       = 0
	YellowCard = 1
	RedCard    = 2
)
