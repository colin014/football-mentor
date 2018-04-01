package model

type DashboardResponse struct {
	CurrentPlace  int        `json:"current_place,omitempty"`
	ClubName      string     `json:"club_name,omitempty"`
	ClubLogoUrl   string     `json:"club_logo_url,omitempty"`
	LeagueName    string     `json:"league_name"`
	LeagueLogoUrl string     `json:"league_logo_url,omitempty"`
	StadiumName   string     `json:"stadium_name,omitempty"`
	WebUrl        string     `json:"web_url,omitempty"`
	NextGame      *GameModel `json:"next_game,omitempty"`
}

type EventType int

const (
	Goal       = 0
	YellowCard = 1
	RedCard    = 2
	OwnGoal    = 3
)
