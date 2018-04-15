package model

type TrainingModel struct {
	BaseModel
	Location     string `json:"location" binding:"required"`
	DayOfTheWeek int    `json:"day_of_the_week" binding:"required"`
	Time         string `json:"time,omitempty" binding:"required"`
	Latitude     string `json:"lat,omitempty"`
	Longitude    string `json:"lon,omitempty"`
}

type UpdateTrainingRequest struct {
	Location     string `json:"location"`
	DayOfTheWeek int    `json:"day_of_the_week"`
	Time         string `json:"time,omitempty"`
	Latitude     string `json:"lat,omitempty"`
	Longitude    string `json:"lat,omitempty"`
}

type CreateTrainingResponse struct {
	Id uint `json:"id" binding:"required"`
}

type TrainingListResponse struct {
	Trainings []TrainingModel `json:"trainings" binding:"required"`
}

func (TrainingModel) TableName() string {
	return "trainings"
}

func (t *TrainingModel) Save() error {
	return db.Save(&t).Error
}

func (t *TrainingModel) Delete() error {
	return db.Delete(&t).Error
}

func (t *TrainingModel) Update(r *UpdateTrainingRequest) error {

	if len(r.Location) != 0 {
		t.Location = r.Location
	}

	//  todo check validity
	if r.DayOfTheWeek != 0 {
		t.DayOfTheWeek = r.DayOfTheWeek
	}

	if len(r.Latitude) != 0 {
		t.Latitude = r.Latitude
	}

	if len(r.Longitude) != 0 {
		t.Longitude = r.Longitude
	}

	if len(r.Time) != 0 {
		t.Time = r.Time
	}

	return t.Save()
}

func GetTrainings() ([]TrainingModel, error) {
	var trainings []TrainingModel
	err := db.Find(&trainings).Error
	return trainings, err
}

func ConvertTrainingsModelToResponse(trainings []TrainingModel) *TrainingListResponse {
	return &TrainingListResponse{Trainings: trainings}
}

func GetTraining(trainingId uint) (*TrainingModel, error) {
	var training TrainingModel
	err := db.Where(TrainingModel{BaseModel: BaseModel{ID: trainingId}}).First(&training).Error
	return &training, err
}
