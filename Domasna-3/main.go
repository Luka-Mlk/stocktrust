package main

import (
	"log"
	"stocktrust/handlers/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})
	app.Static("assets", "./assets")
	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// app.Get("/sse/trading-signals", views.CompanyTradingSignals)
	app.Get("/", views.LandingPage)
	// app.Get("/about-us")
	app.Get("/companies", views.AllCompanies)
	app.Get("/company/:tkr", views.CompanyDetails)

	err := app.Listen(":3030")
	if err != nil {
		log.Println(err)
	}
}
