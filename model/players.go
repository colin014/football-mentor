package model

import "github.com/jinzhu/gorm"

type Player struct {
	gorm.Model          `json:"-"`
	Name         string `json:"name" binding:"required"`
	ImageUrl     string `json:"image_url,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	BirthPlace   string `json:"birth_place,omitempty"`
	Description  string `json:"description,omitempty"`
	JerseyNumber int    `json:"jersey_number,omitempty"`
}

func (p Player) TableName() string {
	return "players"
}
