package main

import (
	"fmt"
	"log"
	"user-service/api/handler"
	"user-service/api/router"
	"user-service/internal/config"
	"user-service/internal/pkg/logs"
	"user-service/internal/service"
	"user-service/storage"
	"user-service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	logs.InitLogger()
	logger := logs.Logger
	log.Println("Server started...")

	cfg := config.Load()
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)

	db, err := postgres.ConnectToPostgres(connStr)
	if err != nil {
		log.Fatal("Error in connect to postgres: ", err)
	}
	defer db.Close()

	queries := storage.New(db)

	r := gin.Default()

	mainHandler := handler.NewMainHandler(service.NewUserService(queries, logger), logger)
	controller := router.NewController(mainHandler, r)
	controller.SetupRoutes()

	log.Fatal(r.Run(":8179"))
}
