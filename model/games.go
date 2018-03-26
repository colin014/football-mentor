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
	ResultId   uint   `gorm:"foreign_key" json:"-"`
	IsHome     bool   `json:"is_home"`
	Type       int    `json:"type"`
	Minute     int    `json:"minute"`
	PlayerName string `json:"player_name"`
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
