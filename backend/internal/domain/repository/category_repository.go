package repository

import (
	"context"
	"todolist-backend/internal/domain/entity"
)

type CategoryRepository interface {
	List(ctx context.Context, limit, offset int) ([]entity.Category, int64, error)
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id uint) (*entity.Category, error)
	GetByName(ctx context.Context, name string) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id uint) error
}