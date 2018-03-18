package api

import (
	"github.com/colin014/football-mentor/config"
	"github.com/colin014/football-mentor/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var db *gorm.DB

func init() {
	logger = config.Logger()
	db = database.GetDB()
}
