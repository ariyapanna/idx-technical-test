package dto

import (
	"time"
	"todolist-backend/internal/usecase"
)

type CreateTodoRequest struct {
	CategoryID  uint   `json:"category_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	Priority    string `json:"priority" binding:"required"`
	DueDate     string `json:"due_date" binding:"omitempty"`
}

func (r *CreateTodoRequest) ToUsecaseInput() usecase.CreateTodoInput {
	return usecase.CreateTodoInput{
		CategoryID:  r.CategoryID,
		Title:       r.Title,
		Description: r.Description,
		Priority:    r.Priority,
		DueDate:     r.DueDate,
	}
}

type UpdateTodoRequest struct {
	CategoryID  uint   `json:"category_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	Priority    string `json:"priority" binding:"required"`
	DueDate     string `json:"due_date" binding:"omitempty"`
	Completed   bool   `json:"completed"`
}

func (r *UpdateTodoRequest) ToUsecaseInput() usecase.UpdateTodoInput {
	return usecase.UpdateTodoInput{
		CategoryID:  r.CategoryID,
		Title:       r.Title,
		Description: r.Description,
		Priority:    r.Priority,
		DueDate:     r.DueDate,
		Completed:   r.Completed,
	}
}

type TodoResponse struct {
	ID          uint              `json:"id"`
	Category    *CategoryResponse `json:"category"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Priority    string            `json:"priority"`
	DueDate     *string           `json:"due_date"`
	Completed   bool              `json:"completed"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func NewTodoResponse(u *usecase.TodoResponse) TodoResponse {
	var category *CategoryResponse
	if u.Category != nil {
		category = &CategoryResponse{
			ID:        u.Category.ID,
			Name:      u.Category.Name,
			Color:     u.Category.Color,
			CreatedAt: u.Category.CreatedAt,
		}
	}

	return TodoResponse{
		ID:          u.ID,
		Category:    category,
		Title:       u.Title,
		Description: u.Description,
		Priority:    u.Priority,
		DueDate:     u.DueDate,
		Completed:   u.Completed,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func NewTodoListResponse(u []usecase.TodoResponse) []TodoResponse {
	res := make([]TodoResponse, len(u))
	for i, v := range u {
		res[i] = NewTodoResponse(&v)
	}
	return res
}
