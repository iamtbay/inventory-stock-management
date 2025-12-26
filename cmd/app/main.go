package main

// @title Inventory & Order Management API
// @version 1.0
// @description This is a sample server for managing stock and orders.
// @host localhost:8080
// @BasePath /

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iamtbay/is-management/internal/adapters/api"
	"github.com/iamtbay/is-management/internal/adapters/postgres"
	"github.com/iamtbay/is-management/internal/config"
	"github.com/iamtbay/is-management/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
	}

	//config
	config := config.LoadConfig()

	conn, err := postgres.NewDB(config.DatabaseURL)
	if err != nil {
		logger.Error("Error connecting to database")
	}
	logger.Info("Connected to database")

	//REPOS
	productRepo := postgres.NewProductRepository(conn)
	orderRepo := postgres.NewOrderRepository(conn)
	logger.Info("Repositories initialized")
	//REPOS END

	//SERVICES
	productSvc := service.NewProductService(productRepo)
	orderSvc := service.NewOrderService(orderRepo, productRepo)
	logger.Info("Services initialized")
	//SERVICES END

	handler := api.NewHTTPHandler(productSvc, orderSvc)
	logger.Info("Handler initialized")

	mux := api.NewRouter(handler)
	logger.Info("Router initialized")

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}

	//SERVER
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Error starting server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down server", "error", err)
	}
	slog.Info("Closing Database connection...")
	conn.Close()
	slog.Info("Server exited properly")
}
