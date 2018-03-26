package model

import (
	"github.com/jinzhu/gorm"
	"github.com/colin014/football-mentor/database"
	"time"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

type ErrorResponse struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
	Error   string `json:"error,omitempty"`
}

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
