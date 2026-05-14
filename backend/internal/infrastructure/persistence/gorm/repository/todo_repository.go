package repository

import (
	"context"
	"fmt"
	"strings"
	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/domain/filter"
	"todolist-backend/internal/infrastructure/persistence/gorm/model"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) List(ctx context.Context, f filter.TodoFilter) ([]entity.Todo, int64, error) {
	var todoModels []model.Todo
	var total int64

	db := r.db.WithContext(ctx).Model(&model.Todo{})

	if f.CategoryID != nil {
		db = db.Where("category_id = ?", *f.CategoryID)
	}
	if f.Completed != nil {
		db = db.Where("completed = ?", *f.Completed)
	}
	if f.Priority != "" {
		db = db.Where("priority = ?", f.Priority)
	}
	if f.Search != "" {
		db = db.Where("title ILIKE ?", "%"+f.Search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	if f.SortBy != "" {
		sortBy = f.SortBy
	}
	sortOrder := "desc"
	if strings.ToLower(f.SortOrder) == "asc" {
		sortOrder = "asc"
	}
	db = db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	limit := 10
	if f.Limit > 0 {
		limit = f.Limit
	}
	offset := 0
	if f.Page > 1 {
		offset = (f.Page - 1) * limit
	}

	err := db.Preload("Category").Limit(limit).Offset(offset).Find(&todoModels).Error
	if err != nil {
		return nil, 0, err
	}

	var todos []entity.Todo
	for _, m := range todoModels {
		todos = append(todos, *r.toEntity(&m))
	}

	return todos, total, nil
}

func (r *todoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	m := r.toModel(todo)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	todo.ID = m.ID
	todo.CreatedAt = m.CreatedAt
	todo.UpdatedAt = m.UpdatedAt
	return nil
}

func (r *todoRepository) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {
	var m model.Todo
	if err := r.db.WithContext(ctx).Preload("Category").First(&m, id).Error; err != nil {
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *todoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	m := r.toModel(todo)
	err := r.db.WithContext(ctx).
		Model(&model.Todo{}).
		Where("id = ?", todo.ID).
		Updates(map[string]interface{}{
			"category_id": todo.CategoryID,
			"title":       todo.Title,
			"description": m.Description,
			"completed":   todo.Completed,
			"priority":    todo.Priority,
			"due_date":    todo.DueDate,
		}).Error
	if err != nil {
		return err
	}
	
	var updated model.Todo
	if err := r.db.WithContext(ctx).Select("updated_at").First(&updated, todo.ID).Error; err == nil {
		todo.UpdatedAt = updated.UpdatedAt
	}

	return nil
}

func (r *todoRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}

func (r *todoRepository) toEntity(m *model.Todo) *entity.Todo {
	var desc string
	if m.Description != nil {
		desc = *m.Description
	}

	var category *entity.Category
	if m.Category != nil {
		category = &entity.Category{
			ID:        m.Category.ID,
			Name:      m.Category.Name,
			Color:     m.Category.Color,
			CreatedAt: m.Category.CreatedAt,
		}
	}

	return &entity.Todo{
		ID:          m.ID,
		CategoryID:  m.CategoryID,
		Category:    category,
		Title:       m.Title,
		Description: desc,
		Completed:   m.Completed,
		Priority:    entity.TodoPriority(m.Priority),
		DueDate:     m.DueDate,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func (r *todoRepository) toModel(e *entity.Todo) *model.Todo {
	var desc *string
	if e.Description != "" {
		desc = &e.Description
	}

	return &model.Todo{
		ID:          e.ID,
		CategoryID:  e.CategoryID,
		Title:       e.Title,
		Description: desc,
		Completed:   e.Completed,
		Priority:    string(e.Priority),
		DueDate:     e.DueDate,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}