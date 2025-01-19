package main

import (
	"log"
	"msemk/handlers"
	"msemk/pkg/queue/dbq"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dbq.DBQueue().Init()

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Post("/records", handlers.CreateRecord)
	v1.Get("/records/top-ten", handlers.GetRecordsTopTen)

	v1.Post("/companies", handlers.CreateCompany)
	v1.Get("/companies", handlers.GetManyCompanies)
	v1.Get("/companies/:tkr", handlers.GetCompany)

	err := app.Listen(os.Getenv("MSEMK_PORT"))
	if err != nil {
		log.Println(err)
	}
}
