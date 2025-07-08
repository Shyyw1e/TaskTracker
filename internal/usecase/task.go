package usecase

import (
	"errors"
	"strings"
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
		normalized := strings.ToLower(tagName)
		var tag db.Tag
		result := database.Where("name = ?", normalized).First(&tag)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				logger.Log.Error("Failed to find tag")
				return nil, result.Error
			}
			tag = db.Tag{Name: normalized}
			
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

	logger.Log.Info("Task created")
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

func GetTasksByTag(database *gorm.DB, userID int64, tagName string) ([]db.Task, error) {
	var tasks []db.Task
	err := database.
		Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Where("tasks.user_id = ? AND tags.name = ?", userID, tagName).
		Preload("Tags").
		Find(&tasks).Error
	if err != nil {
		logger.Log.Errorf("failed to find tasks by tag: %v", err)
		return nil, err
	}

	return tasks, nil
}

func UpdateTaskByID(database *gorm.DB, userID int64, id uint64, req CreateTaskRequest) (*db.Task, error) {
	var task db.Task
	result := database.Preload("Tags").
		Where("user_id = ? AND id = ?", userID, id).First(&task)
	if result.Error != nil {
		logger.Log.Errorf("failed to find task: %v", result.Error)
		return nil, result.Error
	}
	task.Title = req.Title
	task.Description = req.Description
	var tags []db.Tag
	for _, tagName := range req.Tags {
		normalized := strings.ToLower(tagName)
		var tag db.Tag
		result := database.Where("name = ?", normalized).First(&tag)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				logger.Log.Error("Failed to find tag")
				return nil, result.Error
			}
			tag = db.Tag{Name: normalized}
			
			if err := database.Create(&tag).Error; err != nil {
				logger.Log.Errorf("Failed to create tag: %v", err)
				return nil, err
			}
		}
		tags = append(tags, tag)
	}

	err := database.Model(&task).Association("Tags").Replace(tags)
	if err != nil {
		logger.Log.Errorf("failed to replace tags: %v", err)
		return nil, err
	}

	if err := database.Save(&task).Error; err != nil {
		logger.Log.Errorf("failed to update task fields: %v", err)
		return nil, err
	}

	logger.Log.Info("Task updated")
	return &task, nil	
}

func DeleteTaskByID(database *gorm.DB, userID int64, id uint64) error {
	var task db.Task

	result := database.Where("id = ? AND user_id = ?", id, userID).First(&task)
	if err := result.Error; err != nil {
		logger.Log.Errorf("failed to find the task: %v", err)
		return err
	}

	if err := database.Delete(&task).Error; err != nil {
		logger.Log.Errorf("failed to delete the task: %v", err)
		return err
	}

	logger.Log.Info("Task deleted")
	return nil

}