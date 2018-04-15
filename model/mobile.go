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

// todo save to database, if table empty fill it with base menu structure
type ConfigResponse struct {
	Menus []Menu `json:"menus,omitempty"`
}

type Menu struct {
	Title     string     `json:"title,omitempty"`
	Disabled  bool       `json:"disabled,omitempty"`
	MenuItems []MenuItem `json:"menu_items,omitempty"`
}

type MenuItem struct {
	Title    string `json:"title,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type EventType int

const (
	Goal       = 0
	YellowCard = 1
	RedCard    = 2
	OwnGoal    = 3
)
