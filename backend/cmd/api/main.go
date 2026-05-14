package main

import (
	"log"
	"todolist-backend/internal/config"
	"todolist-backend/internal/delivery/http/handler"
	"todolist-backend/internal/delivery/http/router"
	"todolist-backend/internal/infrastructure/database"
	"todolist-backend/internal/infrastructure/persistence/gorm/repository"
	"todolist-backend/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DB.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	// Todo
	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := router.NewRouter(categoryHandler, todoHandler)

	log.Printf("Server running on :%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
