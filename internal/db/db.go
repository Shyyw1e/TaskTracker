package db

import (
	"github.com/Shyyw1e/TaskTracker/config"
	"github.com/Shyyw1e/TaskTracker/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(cfg.DBName))
	if err != nil {
		logger.Log.Error("failed to open db")
		return nil, err
	}
	database.AutoMigrate(&Task{})
	database.AutoMigrate(&Tag{})
	database.AutoMigrate(&TaskTag{})
	return database, nil
}

