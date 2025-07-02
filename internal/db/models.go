package db

import "time"

type Task struct {
	ID 				uint 			`gorm:"primaryKey"`
	Title			string			`gorm:"title"`
	Description		string			`gorm:"description"`
	UserID			int64			`gorm:"user_id"`
	IsDone			bool			`gorm:"is_done"`
	CreatedAt		time.Time		`gorm:"created_at"`
}