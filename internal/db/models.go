package db

import "time"

type Task struct {
	ID 				uint 			`gorm:"primaryKey" json:"id"`
	Title			string			`gorm:"title" json:"title"`
	Description		string			`gorm:"description" json:"description"`
	UserID			int64			`gorm:"user_id" json:"user_id"`
	IsDone			bool			`gorm:"is_done" json:"is_done,omitempty"`
	CreatedAt		time.Time		`gorm:"created_at" json:"created_at,omitempty"`
	Tags 			[]Tag 			`gorm:"many2many:task_tags" json:"tags,omitempty"`
}

type Tag struct {
	ID 				uint			`gorm:"primaryKey"`
	Name			string			`gorm:"unique"`
}

type TaskTag struct {
	TaskID			uint			`gorm:"primaryKey"`
	TagID			uint			`gorm:"primaryKey"`
}