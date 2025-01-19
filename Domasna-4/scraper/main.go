package main

import (
	"log"
	"os"
	"stocktrust/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/scrape", handlers.Scraper)

	err := app.Listen(os.Getenv("SCRAPE_PORT"))
	if err != nil {
		log.Println(err)
	}
}
