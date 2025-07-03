package usecase

import (
	"errors"

	"github.com/Shyyw1e/TaskTracker/internal/db"
	"github.com/Shyyw1e/TaskTracker/pkg/logger"
	"gorm.io/gorm"
)


type CreateTaskRequest struct {
	Title			string		`json:"title"`
	Description		string		`json:"description"`
	UserID			int64		`json:"user_id"`
	Tags			[]string	`json:"tags,omitempty"`		
}


func CreateTask(database *gorm.DB, req CreateTaskRequest) (*db.Task, error){
	var tags []db.Tag
	for _, tagName := range req.Tags {
		var tag db.Tag
		result := database.Where("name = ?", tagName).First(&tag)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				logger.Log.Error("Failed to use database.Where.First func")
				return nil, result.Error
			}
			tag = db.Tag{Name: tagName}
			
			if err := database.Create(&tag).Error; err != nil {
				logger.Log.Errorf("Failed to create tag: %v", err)
				return nil, err
			}
		}
		tags = append(tags, tag)
	}
	newTask := db.Task{
		Title: req.Title,
		Description: req.Description,
		UserID: req.UserID,
		Tags: tags,
	}
	
	if err := database.Create(&newTask).Error; err != nil {
		logger.Log.Errorf("Failed to create task: %v", err)
		return nil, err
	}

	return &newTask, nil

}


func GetAllTasks(database *gorm.DB, userID int64) ([]db.Task, error) {
	var tasks []db.Task
	result := database.Preload("Tags").Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		logger.Log.Errorf("couldn't find tasks: %v", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

