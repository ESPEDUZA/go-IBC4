package main

import (
	"log"
	"os"

	"github.com/ESPEDUZA/go-IBC4/tree/main/fiberTest/handlers"
	"github.com/ESPEDUZA/go-IBC4/tree/main/fiberTest/models"
	"github.com/ESPEDUZA/go-IBC4/tree/main/fiberTest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	models.SetupModels(db)

	app := fiber.New()
	routes.SetupRoutes(app, handlers.NewUserHandler(db))

	log.Fatal(app.Listen(":3000"))
}
