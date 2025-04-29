package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vibhuair20/dsa-master/backend/api/auth"
	"github.com/Vibhuair20/dsa-master/backend/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:api/resolve/v1", routes.Resolve)
	app.Post("/:api/assign/v1", routes.GenerateCalender) // Fixed: was 'assignCodes', should match function name
	app.Get("/:url/admin/v1", routes.adminBoard)
	app.Get("/login", auth.GoogleLogin)            // Fixed: removed parentheses, functions are passed as references
	app.Get("/auth/callback", auth.GoogleCallBack) // Added callback route which was missing
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
