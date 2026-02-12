package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pizza-tracker-go/internal/app/notification"
	"pizza-tracker-go/internal/app/usecase"
	"pizza-tracker-go/internal/infra/config"
	"pizza-tracker-go/internal/infra/delivery/http"
	"pizza-tracker-go/internal/infra/delivery/http/handler"
	"pizza-tracker-go/internal/infra/delivery/http/router"
	"pizza-tracker-go/internal/infra/delivery/http/session"
	"pizza-tracker-go/internal/infra/persistence/gorm"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := gorm.NewDB(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")

	http.RegisterValidators()

	orderRepo := gorm.NewOrderRepository(db)
	userRepo := gorm.NewUserRepository(db)

	orderService := usecase.NewOrderService(orderRepo)
	authService := usecase.NewAuthService(userRepo)
	notifier := notification.NewInMemoryManager()

	h := handler.NewHandler(orderService, authService, notifier)

	engine := gin.Default()

	if err := http.LoadTemplates(engine, "internal/infra/delivery/http/templates/*.tmpl"); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	sessionStore := session.NewStore(db.DB, []byte(cfg.SessionSecretKey))
	router.Setup(engine, h, sessionStore)

	slog.Info("Server starting", "url", "http://localhost:"+cfg.Port)
	_ = engine.Run(":" + cfg.Port)
}
