package model

import "time"

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	CategoryID  *uint `gorm:"index:idx_todos_category_id"`

	Category    *Category `gorm:"foreignKey:CategoryID"`
	Title       string `gorm:"not null;size:255;index:idx_todos_title"`
	Description *string `gorm:"type:text"`
	Completed   bool `gorm:"not null;default:false;index:idx_todos_completed"`
	Priority    string `gorm:"type:todo_priority;not null;index:idx_todos_priority"`
	DueDate     *time.Time

	CreatedAt   time.Time `gorm:"not null;default:now();index:idx_todos_created_at"`
	UpdatedAt   time.Time `gorm:"not null;default:now()"`
}