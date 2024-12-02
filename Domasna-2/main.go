package main

import (
	"log"
	"stocktrust/handlers/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize the app with a template engine
	app := fiber.New(fiber.Config{
		Views: html.New("./templates/html", ".html"),
	})
	app.Static("assets", "./assets")
	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	app.Get("/", views.LandingPage)
	// app.Get("/about-us")
	// v1.Get("/companies")
	app.Get("/company/:tkr", views.GetByTicker)

	err := app.Listen(":3030")
	if err != nil {
		log.Println(err)
	}
}
