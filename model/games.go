package model

import (
	"sort"
	"time"
)

type GameModel struct {
	BaseModel
	IsHome           bool         `json:"is_home"`
	OpponentTeamName string       `json:"opponent_team_name"`
	OpponentTeamLogo string       `json:"opponent_team_logo"`
	Date             string       `json:"date"`
	Time             string       `json:"time"`
	Result           *ResultModel `json:"result,omitempty"`
}

type CreateGameResponse struct {
	Id           uint   `json:"id"`
	OpponentTeam string `json:"opponent_team"`
}

type ResultModel struct {
	GameId   uint `gorm:"primary_key" json:"-"`
	HomeGoal int  `json:"home_goal"`
	AwayGoal int  `json:"away_goal"`
}

type GameListResponse struct {
	Games []GameModel `json:"games" binding:"required"`
}

type EventListResponse struct {
	Events []EventModel `json:"events" binding:"required"`
}

type EventModel struct {
	BaseModel
	GameId     uint   `gorm:"foreign_key" json:"-"`
	IsHome     bool   `json:"is_home" binding:"required"`
	Type       int    `json:"type" binding:"required"`
	Minute     int    `json:"minute" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
}

type CreateEventRequest struct {
	Events []EventModel `json:"events" binding:"required"`
}

func (GameModel) TableName() string {
	return "games"
}

func (ResultModel) TableName() string {
	return "results"
}

func (EventModel) TableName() string {
	return "events"
}

func GetAllGames() ([]GameModel, error) {
	var games []GameModel

	if err := db.Find(&games).Error; err != nil {
		return nil, err
	}

	for i, game := range games {
		games[i].Result = &ResultModel{}
		if err := db.Where(ResultModel{GameId: game.ID}).First(&(*games[i].Result)).Error; err != nil {
			games[i].Result = nil
		}
	}

	return games, nil
}

func ConvertGameModelToResponse(games []GameModel) *GameListResponse {
	return &GameListResponse{Games: games}
}

func ConvertEventModelToResponse(events []EventModel) *EventListResponse {
	return &EventListResponse{Events: events}
}

func (request *CreateEventRequest) SaveEvents(gameId uint) error {

	for _, e := range request.Events {
		e.GameId = gameId
		err := db.Save(&e).Error
		if err != nil {
			return err
		}
	}

	return nil

}

func GetAllEvents(gameId uint) ([]EventModel, error) {
	var events []EventModel
	err := db.Where(EventModel{GameId: gameId}).Find(&events).Error
	return events, err
}

func DeleteEvent(gameId, eventId uint) error {
	return db.Delete(EventModel{GameId: gameId, BaseModel: BaseModel{ID: eventId}}).Error
}

func (r *ResultModel) SaveResult(gameId uint) error {
	r.GameId = gameId
	return db.Save(&r).Error
}

func DeleteResult(gameId uint) error {
	return db.Delete(ResultModel{GameId: gameId}).Error
}

func DeleteGame(gameId uint) error {
	return db.Delete(GameModel{BaseModel: BaseModel{ID: gameId}}).Error
}

func GetNextGame() (*GameModel, error) {
	if games, err := GetAllGames(); err != nil {
		return nil, err
	} else if len(games) == 0 {
		return nil, nil
	} else {

		sort.Slice(games, func(i, j int) bool {
			gameTimeI, _ := time.Parse("20060102", games[i].Date)
			gameTimeJ, _ := time.Parse("20060102", games[j].Date)

			return gameTimeI.Before(gameTimeJ)
		})

		now, _ := time.Parse("20060102", time.Now().Format("20060102"))

		for _, game := range games {
			date, _ := time.Parse("20060102", game.Date)
			if !date.Before(now) {
				return &game, nil
			}
		}

		return &games[0], nil
	}

}
