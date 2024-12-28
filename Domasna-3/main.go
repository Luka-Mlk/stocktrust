package main

import (
	"log"
	"stocktrust/handlers/views"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/indicators"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/k0kubun/pp"
)

func main() {
	dummyrecord := []hrecord.HRecord{
		{
			Id:             "12345",
			Date:           "2024-01-01",
			Ticker:         "AAPL",
			POLT:           5.67,
			Max:            145.32,
			Min:            142.11,
			AvgPrice:       143.56,
			RevenuePercent: 12.45,
			Amount:         15000.0,
			RevenueBEST:    1800000.0,
			RevenueTotal:   2100000.0,
			Currency:       "USD",
		},
		{
			Id:             "12346",
			Date:           "2024-01-02",
			Ticker:         "GOOGL",
			POLT:           3.45,
			Max:            2760.55,
			Min:            2700.12,
			AvgPrice:       2730.83,
			RevenuePercent: 10.78,
			Amount:         10000.0,
			RevenueBEST:    2700000.0,
			RevenueTotal:   3000000.0,
			Currency:       "USD",
		},
		{
			Id:             "12347",
			Date:           "2024-01-03",
			Ticker:         "AMZN",
			POLT:           7.21,
			Max:            3345.21,
			Min:            3270.98,
			AvgPrice:       3308.59,
			RevenuePercent: 15.34,
			Amount:         20000.0,
			RevenueBEST:    4500000.0,
			RevenueTotal:   5200000.0,
			Currency:       "USD",
		},
		{
			Id:             "12348",
			Date:           "2024-01-04",
			Ticker:         "MSFT",
			POLT:           4.12,
			Max:            295.50,
			Min:            290.25,
			AvgPrice:       292.87,
			RevenuePercent: 9.84,
			Amount:         12000.0,
			RevenueBEST:    3400000.0,
			RevenueTotal:   3900000.0,
			Currency:       "USD",
		},
		{
			Id:             "12349",
			Date:           "2024-01-05",
			Ticker:         "TSLA",
			POLT:           6.50,
			Max:            1200.00,
			Min:            1150.60,
			AvgPrice:       1175.30,
			RevenuePercent: 11.22,
			Amount:         8000.0,
			RevenueBEST:    2000000.0,
			RevenueTotal:   2200000.0,
			Currency:       "USD",
		},
		{
			Id:             "12350",
			Date:           "2024-01-06",
			Ticker:         "NVDA",
			POLT:           8.12,
			Max:            930.45,
			Min:            915.78,
			AvgPrice:       922.14,
			RevenuePercent: 14.56,
			Amount:         11000.0,
			RevenueBEST:    2800000.0,
			RevenueTotal:   3100000.0,
			Currency:       "USD",
		},
	}
	pp.Println("indicators:")
	pp.Println(indicators.CalculateIndicators(dummyrecord))
	pp.Println("oscillators:")
	pp.Println(indicators.CalculateOscillators(dummyrecord))
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
