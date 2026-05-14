package usecase

import (
	"context"
	"errors"
	"time"
	
	"todolist-backend/internal/domain/apperror"
	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type CreateCategoryInput struct {
	Name  string
	Color string
}

type UpdateCategoryInput struct {
	Name  string
	Color string
}

type CategoryResponse struct {
	ID        uint
	Name      string
	Color     string
	CreatedAt time.Time
}

type CategoryListResult struct {
	Data       []CategoryResponse
	TotalCount int64
}

type CategoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
	}
}

func (u *CategoryUsecase) Create(ctx context.Context, input CreateCategoryInput) (*CategoryResponse, error) {
	existing, err := u.repo.GetByName(ctx, input.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.NewInternalError(err)
	}
	if existing != nil {
		return nil, apperror.NewConflictError("category with this name already exists")
	}

	category, err := entity.NewCategory(input.Name, input.Color)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(ctx, category)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Color:     category.Color,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (u *CategoryUsecase) List(ctx context.Context, page, limit int) (*CategoryListResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	categories, total, err := u.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	var data []CategoryResponse
	for _, c := range categories {
		data = append(data, CategoryResponse{
			ID:        c.ID,
			Name:      c.Name,
			Color:     c.Color,
			CreatedAt: c.CreatedAt,
		})
	}

	return &CategoryListResult{
		Data:       data,
		TotalCount: total,
	}, nil
}

func (u *CategoryUsecase) GetByID(ctx context.Context, id uint) (*CategoryResponse, error) {
	category, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewNotFoundError("category not found")
		}
		return nil, apperror.NewInternalError(err)
	}

	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Color:     category.Color,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (u *CategoryUsecase) Update(ctx context.Context, id uint, input UpdateCategoryInput) (*CategoryResponse, error) {
	category, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewNotFoundError("category not found")
		}
		return nil, apperror.NewInternalError(err)
	}

	if category.Name != input.Name {
		existing, err := u.repo.GetByName(ctx, input.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewInternalError(err)
		}
		if existing != nil {
			return nil, apperror.NewConflictError("category with this name already exists")
		}
	}

	category.Name = input.Name
	category.Color = input.Color

	if err := category.Validate(); err != nil {
		return nil, err
	}

	err = u.repo.Update(ctx, category)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Color:     category.Color,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (u *CategoryUsecase) Delete(ctx context.Context, id uint) error {
	_, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NewNotFoundError("category not found")
		}
		return apperror.NewInternalError(err)
	}

	err = u.repo.Delete(ctx, id)
	if err != nil {
		return apperror.NewInternalError(err)
	}
	return nil
}