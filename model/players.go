package model

import "github.com/jinzhu/gorm"

type Player struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address" binding:"required"`
	BirthDate  string `json:"birth_date" binding:"required"`
	BirthPlace string `json:"birth_place" binding:"required"`
}

func (p Player) TableName() string {
	return "players"
}
