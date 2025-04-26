package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vibhuair20/dsa-master/backend/api/database"
	"github.com/Vibhuair20/dsa-master/backend/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:api/resolve/v1", routes.Resolve)
	app.Post("/:api/assign/v1", routes.assignCodes)
	app.Get("/:url/admin/v1", routes.adminBoard)

}

type repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	database.CreateClient()

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
