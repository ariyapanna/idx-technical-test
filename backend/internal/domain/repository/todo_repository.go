package repository

import (
	"context"
	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/domain/filter"
)

type TodoRepository interface {
	List(ctx context.Context, filter filter.TodoFilter) ([]entity.Todo, int64, error)
	Create(ctx context.Context, todo *entity.Todo) error
	GetByID(ctx context.Context, id uint) (*entity.Todo, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id uint) error
}