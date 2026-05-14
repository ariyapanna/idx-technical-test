package dto

import (
	"time"
	"todolist-backend/internal/usecase"
)

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (r *CreateCategoryRequest) ToUsecaseInput() usecase.CreateCategoryInput {
	return usecase.CreateCategoryInput{
		Name:  r.Name,
		Color: r.Color,
	}
}

type UpdateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (r *UpdateCategoryRequest) ToUsecaseInput() usecase.UpdateCategoryInput {
	return usecase.UpdateCategoryInput{
		Name:  r.Name,
		Color: r.Color,
	}
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategoryResponse(u *usecase.CategoryResponse) CategoryResponse {
	return CategoryResponse{
		ID:        u.ID,
		Name:      u.Name,
		Color:     u.Color,
		CreatedAt: u.CreatedAt,
	}
}

func NewCategoryListResponse(u []usecase.CategoryResponse) []CategoryResponse {
	res := make([]CategoryResponse, len(u))
	for i, v := range u {
		res[i] = NewCategoryResponse(&v)
	}
	return res
}
