package main

import (
	"log"
	"os"
	"tasks-api/db"
	"tasks-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("env. файл отсутсвует")
	}
}

func main() {
	app := fiber.New()

	service, err := db.NewPostgresService()
	if err != nil {
		log.Fatalf("ошибка при подключении к БД: %v", err)
	}
	defer db.ClosePostgresService(service)

	handlers.RegisterTaskHandlers(app, service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(":3030")
	if err != nil {
		log.Fatal("При запуске сервреа произошла ошибка")
	}
}
