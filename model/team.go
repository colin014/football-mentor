package model

type TeamModel struct {
	BaseModel
	Name string `json:"name" binding:"required"`
}

type CreateTeamResponse struct {
	Id   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TeamListResponse struct {
	Teams []TeamModel `json:"teams" binding:"required"`
}

type UpdateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

func (TeamModel) TableName() string {
	return "teams"
}

func (t *TeamModel) Save() error {
	return db.Save(&t).Error
}

func (t *TeamModel) Delete() error {
	return db.Delete(&t).Error
}

func (t *TeamModel) Update(r *UpdateTeamRequest) error {
	if len(r.Name) != 0 {
		t.Name = r.Name
	}

	return t.Save()
}

func GetTeam(teamId uint) (*TeamModel, error) {
	var team TeamModel
	err := db.Where(StaffModel{BaseModel: BaseModel{ID: teamId}}).First(&team).Error
	return &team, err
}

func GetTeams() ([]TeamModel, error) {
	var t []TeamModel
	err := db.Find(&t).Error
	return t, err
}

func ConvertTeamModelToResponse(teams []TeamModel) *TeamListResponse {
	return &TeamListResponse{Teams: teams}
}
