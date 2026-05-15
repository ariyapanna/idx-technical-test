package usecase_test

import (
	"context"
	"testing"
	"time"

	"todolist-backend/internal/domain/entity"
	"todolist-backend/internal/domain/filter"
	"todolist-backend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockTodoRepository is a manual mock for TodoRepository
type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) List(ctx context.Context, f filter.TodoFilter) ([]entity.Todo, int64, error) {
	args := m.Called(ctx, f)
	return args.Get(0).([]entity.Todo), args.Get(1).(int64), args.Error(2)
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTodoRepository) MarkAsCompleted(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestTodoUsecase_Create(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	t.Run("success create", func(t *testing.T) {
		dueDate := "2026-12-31"
		input := usecase.CreateTodoInput{
			Title:      "Test Task",
			Priority:   "high",
			DueDate:    dueDate,
			CategoryID: 1,
		}

		repo.On("Create", ctx, mock.AnythingOfType("*entity.Todo")).Return(nil).Once()
		
		parsedDate, _ := time.Parse("2006-01-02", dueDate)
		mockEntity := &entity.Todo{
			ID:         1,
			Title:      input.Title,
			Priority:   entity.PriorityHigh,
			DueDate:    &parsedDate,
			CategoryID: &input.CategoryID,
		}
		repo.On("GetByID", ctx, mock.AnythingOfType("uint")).Return(mockEntity, nil).Once()

		res, err := uc.Create(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, input.Title, res.Title)
		assert.Equal(t, "high", res.Priority)
	})

	t.Run("invalid date format", func(t *testing.T) {
		input := usecase.CreateTodoInput{
			Title:   "Invalid Date",
			DueDate: "31-12-2026",
		}
		res, err := uc.Create(ctx, input)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestTodoUsecase_GetByID(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		mockEntity := &entity.Todo{ID: 1, Title: "Existing Task"}
		repo.On("GetByID", ctx, uint(1)).Return(mockEntity, nil).Once()

		res, err := uc.GetByID(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), res.ID)
	})

	t.Run("not found", func(t *testing.T) {
		repo.On("GetByID", ctx, uint(99)).Return(nil, gorm.ErrRecordNotFound).Once()

		res, err := uc.GetByID(ctx, 99)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestTodoUsecase_Update(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	t.Run("success update", func(t *testing.T) {
		id := uint(1)
		existing := &entity.Todo{ID: id, Title: "Old Title"}
		input := usecase.UpdateTodoInput{Title: "New Title", Priority: "low"}

		repo.On("GetByID", ctx, id).Return(existing, nil).Once()
		repo.On("Update", ctx, mock.AnythingOfType("*entity.Todo")).Return(nil).Once()
		
		// Return updated
		updated := &entity.Todo{ID: id, Title: "New Title", Priority: "low"}
		repo.On("GetByID", ctx, id).Return(updated, nil).Once()

		res, err := uc.Update(ctx, id, input)

		assert.NoError(t, err)
		assert.Equal(t, "New Title", res.Title)
	})
}

func TestTodoUsecase_Delete(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		repo.On("GetByID", ctx, uint(1)).Return(&entity.Todo{ID: 1}, nil).Once()
		repo.On("Delete", ctx, uint(1)).Return(nil).Once()

		err := uc.Delete(ctx, 1)
		assert.NoError(t, err)
	})

	t.Run("fail not found", func(t *testing.T) {
		repo.On("GetByID", ctx, uint(2)).Return(nil, gorm.ErrRecordNotFound).Once()

		err := uc.Delete(ctx, 2)
		assert.Error(t, err)
	})
}

func TestTodoUsecase_MarkAsCompleted(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	repo.On("GetByID", ctx, uint(1)).Return(&entity.Todo{ID: 1, Completed: false}, nil).Once()
	repo.On("MarkAsCompleted", ctx, uint(1)).Return(nil).Once()
	repo.On("GetByID", ctx, uint(1)).Return(&entity.Todo{ID: 1, Completed: true}, nil).Once()

	res, err := uc.MarkAsCompleted(ctx, 1)

	assert.NoError(t, err)
	assert.True(t, res.Completed)
}

func TestTodoUsecase_List(t *testing.T) {
	repo := new(MockTodoRepository)
	uc := usecase.NewTodoUsecase(repo)
	ctx := context.Background()

	mockTodos := []entity.Todo{
		{ID: 1, Title: "Task 1"},
		{ID: 2, Title: "Task 2"},
	}

	f := filter.TodoFilter{Page: 1, Limit: 10}
	repo.On("List", ctx, f).Return(mockTodos, int64(2), nil).Once()

	res, err := uc.List(ctx, f)

	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
	assert.Equal(t, int64(2), res.TotalCount)
}
