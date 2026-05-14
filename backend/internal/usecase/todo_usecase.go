package usecase

import (
	"context"
	"errors"
	"time"
	"todolist-backend/internal/domain/apperror"
	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/domain/filter"
	"todolist-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type CreateTodoInput struct {
	CategoryID  uint
	Title       string
	Description string
	Priority    string
	DueDate     string
}

type UpdateTodoInput struct {
	CategoryID  uint
	Title       string
	Description string
	Priority    string
	DueDate     string
	Completed   bool
}

type TodoResponse struct {
	ID          uint
	Category    *CategoryResponse
	Title       string
	Description string
	Priority    string
	DueDate     *string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoListResult struct {
	Data       []TodoResponse
	TotalCount int64
}

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (u *TodoUsecase) Create(ctx context.Context, input CreateTodoInput) (*TodoResponse, error) {
	var dueDate *time.Time
	if input.DueDate != "" {
		parsed, err := time.Parse("2006-01-02", input.DueDate)
		if err != nil {
			return nil, apperror.NewValidationError("invalid due date format, use YYYY-MM-DD")
		}
		dueDate = &parsed
	}

	todo, err := entity.NewTodo(&input.CategoryID, input.Title, input.Description, entity.TodoPriority(input.Priority), dueDate)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, todo); err != nil {
		return nil, apperror.NewInternalError(err)
	}

	fullTodo, err := u.repo.GetByID(ctx, todo.ID)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	return u.mapToResponse(fullTodo), nil
}

func (u *TodoUsecase) List(ctx context.Context, f filter.TodoFilter) (*TodoListResult, error) {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}

	todos, total, err := u.repo.List(ctx, f)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	var data []TodoResponse
	for _, t := range todos {
		data = append(data, *u.mapToResponse(&t))
	}

	return &TodoListResult{
		Data:       data,
		TotalCount: total,
	}, nil
}

func (u *TodoUsecase) GetByID(ctx context.Context, id uint) (*TodoResponse, error) {
	todo, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewNotFoundError("todo not found")
		}
		return nil, apperror.NewInternalError(err)
	}

	return u.mapToResponse(todo), nil
}

func (u *TodoUsecase) Update(ctx context.Context, id uint, input UpdateTodoInput) (*TodoResponse, error) {
	todo, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewNotFoundError("todo not found")
		}
		return nil, apperror.NewInternalError(err)
	}

	var dueDate *time.Time
	if input.DueDate != "" {
		parsed, err := time.Parse("2006-01-02", input.DueDate)
		if err != nil {
			return nil, apperror.NewValidationError("invalid due date format, use YYYY-MM-DD")
		}
		dueDate = &parsed
	}

	todo.CategoryID = &input.CategoryID
	todo.Title = input.Title
	todo.Description = input.Description
	todo.Priority = entity.TodoPriority(input.Priority)
	todo.DueDate = dueDate
	todo.Completed = input.Completed

	if err := todo.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, todo); err != nil {
		return nil, apperror.NewInternalError(err)
	}

	fullTodo, err := u.repo.GetByID(ctx, todo.ID)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	return u.mapToResponse(fullTodo), nil
}

func (u *TodoUsecase) Delete(ctx context.Context, id uint) error {
	_, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NewNotFoundError("todo not found")
		}
		return apperror.NewInternalError(err)
	}

	if err := u.repo.Delete(ctx, id); err != nil {
		return apperror.NewInternalError(err)
	}

	return nil
}

func (u *TodoUsecase) MarkAsCompleted(ctx context.Context, id uint) (*TodoResponse, error) {
	todo, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewNotFoundError("todo not found")
		}
		return nil, apperror.NewInternalError(err)
	}

	if err := u.repo.MarkAsCompleted(ctx, todo.ID); err != nil {
		return nil, apperror.NewInternalError(err)
	}

	fullTodo, err := u.repo.GetByID(ctx, todo.ID)
	if err != nil {
		return nil, apperror.NewInternalError(err)
	}

	return u.mapToResponse(fullTodo), nil
}

func (u *TodoUsecase) mapToResponse(todo *entity.Todo) *TodoResponse {
	var dueDateStr *string
	if todo.DueDate != nil {
		s := todo.DueDate.Format("2006-01-02")
		dueDateStr = &s
	}

	var category *CategoryResponse
	if todo.Category != nil {
		category = &CategoryResponse{
			ID:        todo.Category.ID,
			Name:      todo.Category.Name,
			Color:     todo.Category.Color,
			CreatedAt: todo.Category.CreatedAt,
		}
	}

	return &TodoResponse{
		ID:          todo.ID,
		Category:    category,
		Title:       todo.Title,
		Description: todo.Description,
		Priority:    string(todo.Priority),
		DueDate:     dueDateStr,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
