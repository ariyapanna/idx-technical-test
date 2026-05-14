package handler

import (
	"math"
	"net/http"
	"strconv"

	http_helper "todolist-backend/internal/delivery/http"
	"todolist-backend/internal/delivery/http/dto"
	"todolist-backend/internal/domain/filter"
	"todolist-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	usecase *usecase.TodoUsecase
}

func NewTodoHandler(u *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase: u,
	}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	response, err := h.usecase.Create(c.Request.Context(), req.ToUsecaseInput())
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create todo", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusCreated, "Todo created successfully", dto.NewTodoResponse(response))
}

func (h *TodoHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	sortBy := c.Query("sort_by")
	sortOrder := c.Query("sort_order")

	f := filter.TodoFilter{
		Page:       page,
		Limit:      limit,
		Search:     search,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
	}

	result, err := h.usecase.List(c.Request.Context(), f)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch todos", err)
		return
	}

	totalPages := int(math.Ceil(float64(result.TotalCount) / float64(limit)))

	meta := http_helper.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		TotalItems:  result.TotalCount,
		TotalPages:  totalPages,
	}

	http_helper.PaginationResponse(
		c,
		http.StatusOK,
		"Todos fetched successfully",
		dto.NewTodoListResponse(result.Data),
		meta,
	)
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid todo ID", err)
		return
	}

	response, err := h.usecase.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusNotFound, "Todo not found", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Todo fetched successfully", dto.NewTodoResponse(response))
}

func (h *TodoHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid todo ID", err)
		return
	}

	var req dto.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	response, err := h.usecase.Update(c.Request.Context(), uint(id), req.ToUsecaseInput())
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to update todo", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Todo updated successfully", dto.NewTodoResponse(response))
}

func (h *TodoHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid todo ID", err)
		return
	}

	err = h.usecase.Delete(c.Request.Context(), uint(id))
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete todo", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Todo deleted successfully", nil)
}

func (h *TodoHandler) Complete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid todo ID", err)
		return
	}

	response, err := h.usecase.MarkAsCompleted(c.Request.Context(), uint(id))
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to complete todo", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Todo status updated successfully", dto.NewTodoResponse(response))
}
