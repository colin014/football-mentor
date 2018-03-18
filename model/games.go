package model

import "github.com/jinzhu/gorm"

type Game struct {
	gorm.Model
	IsHome           bool    `json:"is_home"`
	OpponentTeamName string  `json:"opponent_team_name"`
	OpponentTeamLogo string  `json:"opponent_team_logo"`
	Date             string  `json:"date"`
	Time             string  `json:"time"`
	Result           *Result `json:"result,omitempty"`
}

type Result struct {
	GameId   uint    `gorm:"primary_key"`
	HomeGoal int     `json:"home_goal"`
	AwayGoal int     `json:"away_goal"`
	Events   []Event `json:"events"`
}

type Event struct {
	ResultId   uint   `gorm:"primary_key"`
	IsHome     bool   `json:"is_home"`
	Type       int    `json:"type"`
	Minute     int    `json:"minute"`
	PlayerName string `json:"player_name"`
}

func (Game) TableName() string {
	return "games"
}

func (Result) TableName() string {
	return "results"
}

func (Event) TableName() string {
	return "events"
}
