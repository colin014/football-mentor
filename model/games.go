package model

import (
	"sort"
	"time"
)

type GameModel struct {
	BaseModel
	IsHome           bool         `json:"is_home"`
	OpponentTeamName string       `json:"opponent_team_name,omitempty"`
	OpponentTeamLogo string       `json:"opponent_team_logo,omitempty"`
	Date             string       `json:"date,omitempty"`
	Time             string       `json:"time,omitempty"`
	Result           *ResultModel `json:"result,omitempty"`
}

type UpdateGameRequest struct {
	IsHome           *bool  `json:"is_home,omitempty"`
	OpponentTeamName string `json:"opponent_team_name,omitempty"`
	OpponentTeamLogo string `json:"opponent_team_logo,omitempty"`
	Date             string `json:"date,omitempty"`
	Time             string `json:"time,omitempty"`
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

type UpdateResultRequest struct {
	HomeGoal *int `json:"home_goal,omitempty"`
	AwayGoal *int `json:"away_goal,omitempty"`
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

type UpdateEventModel struct {
	IsHome     *bool  `json:"is_home,omitempty"`
	Type       *int   `json:"type,omitempty"`
	Minute     *int   `json:"minute,omitempty"`
	PlayerName string `json:"player_name,omitempty"`
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

func (g *GameModel) Update(r *UpdateGameRequest) error {
	if r.IsHome != nil {
		g.IsHome = *r.IsHome
	}

	if len(r.OpponentTeamName) != 0 {
		g.OpponentTeamName = r.OpponentTeamName
	}

	if len(r.OpponentTeamLogo) != 0 {
		g.OpponentTeamLogo = r.OpponentTeamLogo
	}

	if len(r.Date) != 0 {
		g.Date = r.Date
	}

	if len(r.Time) != 0 {
		g.Time = r.Time
	}

	return db.Save(&g).Error
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
	// sort before return
	sortGames(games)
	return &GameListResponse{Games: games}
}

func ConvertEventModelToResponse(events []EventModel) *EventListResponse {
	sortEvents(events)
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

func (r *ResultModel) Update(req *UpdateResultRequest) error {

	if req.HomeGoal != nil {
		r.HomeGoal = *req.HomeGoal
	}

	if req.AwayGoal != nil {
		r.AwayGoal = *req.AwayGoal
	}

	return db.Save(&r).Error
}

func (e *EventModel) Update(r *UpdateEventModel) error {
	if r.IsHome != nil {
		e.IsHome = *r.IsHome
	}

	if r.Type != nil {
		e.Type = *r.Type
	}

	if r.Minute != nil {
		e.Minute = *r.Minute
	}

	if len(r.PlayerName) != 0 {
		e.PlayerName = r.PlayerName
	}

	return db.Save(&e).Error
}

func GetResult(gameId uint) (*ResultModel, error) {
	var result ResultModel
	if err := db.Where(ResultModel{GameId: gameId}).First(&result).Error; err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func GetEvent(gameId, eventId uint) (EventModel, error) {
	var event EventModel
	err := db.Where(EventModel{BaseModel: BaseModel{ID: eventId}, GameId: gameId}).First(&event).Error
	return event, err
}

func GetAllEvents(gameId uint) ([]EventModel, error) {
	var events []EventModel
	err := db.Where(EventModel{GameId: gameId}).Find(&events).Error
	return events, err
}

func DeleteEvents(gameId uint) error {
	return db.Delete(EventModel{GameId: gameId}).Error
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

func GetGame(gameId uint) (*GameModel, error) {
	var game GameModel
	err := db.Where(GameModel{BaseModel: BaseModel{ID: gameId}}).First(&game).Error
	return &game, err
}

func (g *GameModel) DeleteGame() error {
	err := DeleteResult(g.ID)
	if err != nil {
		return err
	}

	err = DeleteEvents(g.ID)
	if err != nil {
		return err
	}

	return db.Delete(&g).Error
}

func GetNextGame() (*GameModel, error) {
	if games, err := GetAllGames(); err != nil {
		return nil, err
	} else if len(games) == 0 {
		return nil, nil
	} else {

		sortGames(games)

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

func sortGames(games []GameModel) {
	sort.Slice(games, func(i, j int) bool {
		gameTimeI, _ := time.Parse("20060102", games[i].Date)
		gameTimeJ, _ := time.Parse("20060102", games[j].Date)

		return gameTimeI.Before(gameTimeJ)
	})
}

func sortEvents(events []EventModel) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Minute < events[j].Minute
	})
}
