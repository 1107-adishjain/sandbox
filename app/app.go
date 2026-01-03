package app

import (
	"github.com/1107-adishjain/sandbox/config"
	"gorm.io/gorm"
)

type Application struct {
	Cfg *config.Config
	DB  *gorm.DB
}
