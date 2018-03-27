package model

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

type Event struct {
	BaseModel
	GameId     uint   `gorm:"foreign_key" json:"-"`
	IsHome     bool   `json:"is_home" binding:"required"`
	Type       int    `json:"type" binding:"required"`
	Minute     int    `json:"minute" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
}

type CreateEventRequest struct {
	Events []Event `json:"events" binding:"required"`
}

func (GameModel) TableName() string {
	return "games"
}

func (ResultModel) TableName() string {
	return "results"
}

func (Event) TableName() string {
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

func GetAllEvents(gameId uint) ([]Event, error) {
	var events []Event
	err := db.Where(Event{GameId: gameId}).Find(&events).Error
	return events, err
}

func DeleteEvent(gameId, eventId uint) error {
	return db.Delete(Event{GameId: gameId, BaseModel: BaseModel{ID: eventId}}).Error
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
