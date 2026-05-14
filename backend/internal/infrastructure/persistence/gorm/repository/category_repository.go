package repository

import (
	"context"
	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/infrastructure/persistence/gorm/model"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) List(ctx context.Context, limit, offset int) ([]entity.Category, int64, error) {
	var categories []model.Category
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Category{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Limit(limit).Offset(offset).Order("created_at desc").Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	var result []entity.Category
	for _, c := range categories {
		result = append(result, entity.Category{
			ID:        c.ID,
			Name:      c.Name,
			Color:     c.Color,
			CreatedAt: c.CreatedAt,
		})
	}

	return result, total, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	categoryModel := model.Category{
		Name:  category.Name,
		Color: category.Color,
	}

	err := r.db.WithContext(ctx).Create(&categoryModel).Error
	if err != nil {
		return err
	}

	category.ID = categoryModel.ID
	category.CreatedAt = categoryModel.CreatedAt

	return nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id uint) (*entity.Category, error) {
	var categoryModel model.Category

	err := r.db.WithContext(ctx).First(&categoryModel, id).Error
	if err != nil {
		return nil, err
	}

	return &entity.Category{
		ID:        categoryModel.ID,
		Name:      categoryModel.Name,
		Color:     categoryModel.Color,
		CreatedAt: categoryModel.CreatedAt,
	}, nil
}

func (r *categoryRepository) GetByName(ctx context.Context, name string) (*entity.Category, error) {
	var categoryModel model.Category

	err := r.db.WithContext(ctx).Where("name = ?", name).First(&categoryModel).Error
	if err != nil {
		return nil, err
	}

	return &entity.Category{
		ID:        categoryModel.ID,
		Name:      categoryModel.Name,
		Color:     categoryModel.Color,
		CreatedAt: categoryModel.CreatedAt,
	}, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	err := r.db.WithContext(ctx).
		Model(&model.Category{}).
		Where("id = ?", category.ID).
		Updates(map[string]any{
			"name":  category.Name,
			"color": category.Color,
		}).Error

	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).
		Delete(&model.Category{}, id).Error

	return err
}