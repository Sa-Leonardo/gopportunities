package handler

import (
	"github.com/Sa-Leonardo/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {

	logger = config.GETLogger("handler")
	db = config.GetSQLite()
	_ = db
}
