package model

type PlayerModel struct {
	BaseModel
	Name         string `json:"name" binding:"required"`
	ImageUrl     string `json:"image_url,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	BirthPlace   string `json:"birth_place,omitempty"`
	Description  string `json:"description,omitempty"`
	JerseyNumber int    `json:"jersey_number,omitempty"`
}

type UpdatePlayerRequest struct {
	Name         string `json:"name,omitempty"`
	ImageUrl     string `json:"image_url,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	BirthPlace   string `json:"birth_place,omitempty"`
	Description  string `json:"description,omitempty"`
	JerseyNumber int    `json:"jersey_number,omitempty"`
}

type PlayerListResponse struct {
	Players []PlayerModel `json:"players" binding:"required"`
}

type CreatePlayerResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (p PlayerModel) TableName() string {
	return "players"
}

func (p *PlayerModel) Save() error {
	return db.Save(&p).Error
}

func (p *PlayerModel) Delete() error {
	return db.Delete(&p).Error
}

func (p *PlayerModel) Update(request *UpdatePlayerRequest) error {

	if len(request.Name) != 0 {
		p.Name = request.Name
	}

	if len(request.ImageUrl) != 0 {
		p.ImageUrl = request.ImageUrl
	}

	if len(request.BirthDate) != 0 {
		p.BirthDate = request.BirthDate
	}

	if len(request.BirthPlace) != 0 {
		p.BirthPlace = request.BirthPlace
	}

	if len(request.Description) != 0 {
		p.Description = request.Description
	}

	if request.JerseyNumber != 0 {
		p.JerseyNumber = request.JerseyNumber
	}

	return db.Save(&p).Error
}

func GetPlayer(playerId uint) (*PlayerModel, error) {
	var player PlayerModel
	err := db.Where(PlayerModel{BaseModel: BaseModel{ID: playerId}}).First(&player).Error
	return &player, err
}

func GetAllPlayer() ([]PlayerModel, error) {
	var players []PlayerModel

	if err := db.Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func ConvertPlayerModelToResponse(players []PlayerModel) *PlayerListResponse {
	return &PlayerListResponse{Players: players}
}

func GetPlayerById(id uint, output *PlayerModel) error {
	return db.Where(PlayerModel{BaseModel: BaseModel{ID: id}}).First(&output).Error
}
