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

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	r := router.NewRouter(categoryHandler)

	log.Printf("Server running on :%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
