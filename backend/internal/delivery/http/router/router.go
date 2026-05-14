package router

import (
	"todolist-backend/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(categoryHandler *handler.CategoryHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.Create)
			categories.GET("", categoryHandler.List)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}
	}

	return r
}
