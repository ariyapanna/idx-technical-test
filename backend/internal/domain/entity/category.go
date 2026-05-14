package entity

import (
	"strings"
	"time"
	"todolist-backend/internal/domain/apperror"
)

type Category struct {
	ID    uint
	Name  string
	Color string

	CreatedAt time.Time
}

func (c *Category) ValidateName() error {
	trimmedName := strings.TrimSpace(c.Name)

	if trimmedName == "" {
		return apperror.NewValidationError("name cannot be empty")
	}

	if len(trimmedName) < 3 {
		return apperror.NewValidationError("name must be at least 3 characters")
	}

	if len(trimmedName) > 255 {
		return apperror.NewValidationError("name cannot exceed 255 characters")
	}

	return nil
}

func (c *Category) ValidateColor() error {
	trimmed := strings.TrimSpace(c.Color)

	if trimmed == "" {
		return apperror.NewValidationError("color cannot be empty")
	}

	return nil
}

func (c *Category) Validate() error {
	if err := c.ValidateName(); err != nil {
		return err
	}

	if err := c.ValidateColor(); err != nil {
		return err
	}

	return nil
}

func NewCategory(name, color string) (*Category, error) {
	c := &Category{
		Name:  name,
		Color: color,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}
