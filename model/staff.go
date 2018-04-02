package model

type StaffModel struct {
	BaseModel
	Name        string `json:"name" binding:"required"`
	ImageUrl    string `json:"image_url,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description" binding:"required"`
}

type UpdateStaffRequest struct {
	Name        string `json:"name,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description"`
}

type StaffListResponse struct {
	Staff []StaffModel `json:"staff" binding:"required"`
}

type CreateStaffResponse struct {
	Id   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (StaffModel) TableName() string {
	return "staff"
}

func (s *StaffModel) Update(r *UpdateStaffRequest) error {
	if len(r.Name) != 0 {
		s.Name = r.Name
	}

	if len(r.Phone) != 0 {
		s.Phone = r.Phone
	}

	if len(r.Email) != 0 {
		s.Email = r.Email
	}

	if len(r.Description) != 0 {
		s.Description = r.Description
	}

	if len(r.ImageUrl) != 0 {
		s.ImageUrl = r.ImageUrl
	}

	return db.Save(&s).Error
}

func (s *StaffModel) Save() error {
	return db.Save(&s).Error
}

func (s *StaffModel) Delete() error {
	return db.Delete(&s).Error
}

func GetStaffMember(staffId uint) (StaffModel, error) {
	var staff StaffModel
	err := db.Where(StaffModel{BaseModel: BaseModel{ID: staffId}}).First(&staff).Error
	return staff, err
}

func GetStaff() ([]StaffModel, error) {
	var staff []StaffModel
	err := db.Find(&staff).Error
	return staff, err
}

func ConvertStaffModelToResponse(staff []StaffModel) *StaffListResponse {
	return &StaffListResponse{Staff: staff}
}
