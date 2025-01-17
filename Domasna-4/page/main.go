package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./template", "html"),
	})
	app.Static("assets", "./assets")

	app.Get("/")
	app.Get("/companies")
	app.Get("/company/:tkr")
	app.Get("/about-us")

	err := app.Listen(os.Getenv("PAGE_PORT"))
	if err != nil {
		log.Println(err)
	}
}
