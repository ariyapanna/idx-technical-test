package model

import "time"

type Category struct {
	ID        uint 		`gorm:"primaryKey"`
	Name      string	`gorm:"unique;not null"`
	Color     string	`gorm:"not null"`
	CreatedAt time.Time

	// Todos []Todo `gorm:"foreignKey:CategoryID"`
}