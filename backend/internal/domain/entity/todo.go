package entity

import (
	"strings"
	"time"
	"todolist-backend/internal/domain/apperror"
)

type TodoPriority string

const (
	PriorityHigh   TodoPriority = "high"
	PriorityMedium TodoPriority = "medium"
	PriorityLow    TodoPriority = "low"
)

type Todo struct {
	ID         uint
	CategoryID *uint
	Category   *Category

	Title       string
	Description string
	Completed   bool
	Priority    TodoPriority
	DueDate     *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) ValidateTitle() error {
	trimmedTitle := strings.TrimSpace(t.Title)
	if trimmedTitle == "" {
		return apperror.NewValidationError("title cannot be empty")
	}

	if len(trimmedTitle) < 3 {
		return apperror.NewValidationError("title must be at least 3 characters")
	}

	if len(trimmedTitle) > 255 {
		return apperror.NewValidationError("title cannot exceed 255 characters")
	}

	return nil
}

func (t *Todo) ValidatePriority() error {
	switch t.Priority {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return nil
	default:
		return apperror.NewValidationError("invalid priority value")
	}
}

func (t *Todo) ValidateDueDate() error {
	if t.DueDate != nil && t.DueDate.Before(time.Now().Truncate(time.Hour*24)) {
		return apperror.NewValidationError("due date cannot be in the past")
	}
	return nil
}

func (t *Todo) Validate() error {
	if err := t.ValidateTitle(); err != nil {
		return err
	}

	if err := t.ValidatePriority(); err != nil {
		return err
	}

	if err := t.ValidateDueDate(); err != nil {
		return err
	}

	return nil
}

func NewTodo(categoryID *uint, title, description string, priority TodoPriority, dueDate *time.Time) (*Todo, error) {
	todo := &Todo{
		CategoryID:  categoryID,
		Title:       title,
		Description: description,
		Priority:    priority,
		DueDate:     dueDate,
		Completed:   false,
	}

	if err := todo.Validate(); err != nil {
		return nil, err
	}

	return todo, nil
}