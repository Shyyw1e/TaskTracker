package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Shyyw1e/TaskTracker/internal/usecase"
	"github.com/Shyyw1e/TaskTracker/pkg/logger"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	req := usecase.CreateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		logger.Log.Errorf("failed to decode json: %v", err)
		return
	}
	defer r.Body.Close()
	if req.UserID <= 0 {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		logger.Log.Errorf("invalid user_id: %d", req.UserID)
		return
	}

	task, err := usecase.CreateTask(database, req)
	if err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		logger.Log.Errorf("failed to create task: %v", err)
		return 
	}
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
    	http.Error(w, "user_id is required", http.StatusBadRequest)
    	return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		logger.Log.Errorf("failed to convert userID to int64: %v", err)
		return
	}
	
	if userID <= 0 {
    	http.Error(w, "invalid user_id", http.StatusBadRequest)
    	return
	}

	tasks, err := usecase.GetAllTasks(database, userID)
	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		logger.Log.Errorf("failed to get tasks: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)
}

func GetTasksByTagHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
    	http.Error(w, "user_id is required", http.StatusBadRequest)
    	return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		logger.Log.Errorf("failed to conver userID to int64: %v", err)
		return
	}
	
	if userID <= 0 {
    	http.Error(w, "invalid user_id", http.StatusBadRequest)
    	return
	}

	tagName := r.URL.Query().Get("tag")
	if tagName == "" {
		http.Error(w, "invalid tag (empty)", http.StatusBadRequest)
		logger.Log.Error("Empty tag")
		return
	}

	tasks, err := usecase.GetTasksByTag(database, userID, tagName)
	if err != nil {
		http.Error(w, "failed to get tasks by tag", http.StatusInternalServerError)
		logger.Log.Errorf("failed to get tasks by tag: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)	
}

func UpdateTaskByIDHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		logger.Log.Errorf("invalid id: %v", err)
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
    	http.Error(w, "user_id is required", http.StatusBadRequest)
    	logger.Log.Errorf("user_id is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		logger.Log.Errorf("failed to conver userID to int64: %v", err)
		return
	}
	
	if userID <= 0 {
    	http.Error(w, "invalid user_id", http.StatusBadRequest)
    	return
	}

	req := usecase.CreateTaskRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		logger.Log.Errorf("failed to decode json: %v", err)
		return
	}
	defer r.Body.Close()

	task, err := usecase.UpdateTaskByID(database, userID, idint, req)
	if err != nil {
		http.Error(w, "failed to update the task", http.StatusInternalServerError)
		logger.Log.Errorf("failed to update the task: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskByIDHandler(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		logger.Log.Errorf("invalid id: %v", err)
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
    	http.Error(w, "user_id is required", http.StatusBadRequest)
    	logger.Log.Errorf("user_id is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		logger.Log.Errorf("failed to conver userID to int64: %v", err)
		return
	}
	
	if userID <= 0 {
    	http.Error(w, "invalid user_id", http.StatusBadRequest)
    	return
	}

	if err := usecase.DeleteTaskByID(database, userID, idint); err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}