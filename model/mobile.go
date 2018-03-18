package model

type GetMobileData struct {
	TeamName       string `json:"team_name"`
	TeamLogoUrl    string `json:"team_logo_url"`
	LeagueLogoUrl  string `json:"league_logo_url"`
	CurrentPlace   int    `json:"current_place"`
	StadiumName    string `json:"stadium_name"`
	LeagueName     string `json:"league_name"`
	FacebookPageId string `json:"facebook_page_id"`
	Website        string `json:"website"`
	NextGame       Game   `json:"next_game"`
	Teams          []Team `json:"teams"`
	Games          []Game `json:"games"`
}

type Game struct {
	IsHome           bool    `json:"is_home"`
	OpponentTeamName string  `json:"opponent_team_name"`
	OpponentTeamLogo string  `json:"opponent_team_logo"`
	Date             string  `json:"date"`
	Time             string  `json:"time"`
	Result           *Result `json:"result,omitempty"`
}

type Result struct {
	HomeGoal int     `json:"home_goal"`
	AwayGoal int     `json:"away_goal"`
	Events   []Event `json:"events"`
}

type Event struct {
	IsHome     bool   `json:"is_home"`
	Type       int    `json:"type"`
	Minute     int    `json:"minute"`
	PlayerName string `json:"player_name"`
}

type Team struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type EventType int

const (
	Goal       = 0
	YellowCard = 1
	RedCard    = 2
)
