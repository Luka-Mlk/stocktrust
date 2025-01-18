package main

import (
	"log"
	"os"
	"pages/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})
	app.Static("assets", "./assets")

	app.Get("/", handlers.LandingPage)
	app.Get("/companies", handlers.ManyCompanies)
	app.Get("/company/:tkr", handlers.Company)
	// app.Get("/about-us")

	err := app.Listen(os.Getenv("PAGE_PORT"))
	if err != nil {
		log.Println(err)
	}
}
