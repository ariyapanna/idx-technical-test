package handler

import (
	"math"
	"net/http"
	"strconv"
	http_helper "todolist-backend/internal/delivery/http"
	"todolist-backend/internal/delivery/http/dto"
	"todolist-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(u *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: u,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	response, err := h.usecase.Create(c.Request.Context(), req.ToUsecaseInput())
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusCreated, "Category created successfully", dto.NewCategoryResponse(response))
}

func (h *CategoryHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	result, err := h.usecase.List(c.Request.Context(), page, limit)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch categories", err)
		return
	}

	totalPages := int(math.Ceil(float64(result.TotalCount) / float64(limit)))

	meta := http_helper.PaginationMeta{
		CurrentPage:  page,
		PerPage: limit,
		TotalItems:   result.TotalCount,
		TotalPages:   totalPages,
	}

	http_helper.PaginationResponse(
		c,
		http.StatusOK,
		"Categories fetched successfully",
		dto.NewCategoryListResponse(result.Data),
		meta,
	)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	response, err := h.usecase.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusNotFound, "Category not found", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Category fetched successfully", dto.NewCategoryResponse(response))
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	response, err := h.usecase.Update(c.Request.Context(), uint(id), req.ToUsecaseInput())
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to update category", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Category updated successfully", dto.NewCategoryResponse(response))
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	err = h.usecase.Delete(c.Request.Context(), uint(id))
	if err != nil {
		http_helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete category", err)
		return
	}

	http_helper.SuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}
