package views

import (
	"fmt"
	"log"
	"stocktrust/pkg/company"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/indicators"
	rparser "stocktrust/pkg/strings/parser/records"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LandingPage(c *fiber.Ctx) error {
	records, err := hrecord.GetTopTen()
	if err != nil {
		log.Println("error getting landing page: ", err)
		return c.Render("views/404", nil)
	}
	recordsDemo := []hrecord.HRecord{
		{
			Id:             "H12345",
			Date:           "2024-12-23",
			Ticker:         "AAPL",
			POLT:           0.35,
			Max:            185.50,
			Min:            172.10,
			AvgPrice:       178.80,
			RevenuePercent: 12.5,
			Amount:         100000,
			RevenueBEST:    1500000,
			RevenueTotal:   1800000,
			Currency:       "USD",
		},
		{
			Id:             "H67890",
			Date:           "2024-12-22",
			Ticker:         "GOOGL",
			POLT:           0.42,
			Max:            2950.30,
			Min:            2850.00,
			AvgPrice:       2900.15,
			RevenuePercent: 9.8,
			Amount:         50000,
			RevenueBEST:    700000,
			RevenueTotal:   800000,
			Currency:       "USD",
		},
		{
			Id:             "H11223",
			Date:           "2024-12-21",
			Ticker:         "TSLA",
			POLT:           0.28,
			Max:            750.20,
			Min:            720.45,
			AvgPrice:       735.32,
			RevenuePercent: 15.0,
			Amount:         150000,
			RevenueBEST:    2250000,
			RevenueTotal:   2600000,
			Currency:       "USD",
		},
		{
			Id:             "H33456",
			Date:           "2024-12-20",
			Ticker:         "MSFT",
			POLT:           0.32,
			Max:            310.25,
			Min:            295.75,
			AvgPrice:       302.00,
			RevenuePercent: 11.0,
			Amount:         120000,
			RevenueBEST:    1500000,
			RevenueTotal:   1700000,
			Currency:       "USD",
		},
		{
			Id:             "H77889",
			Date:           "2024-12-19",
			Ticker:         "AMZN",
			POLT:           0.25,
			Max:            3400.00,
			Min:            3300.50,
			AvgPrice:       3350.25,
			RevenuePercent: 10.5,
			Amount:         70000,
			RevenueBEST:    700000,
			RevenueTotal:   800000,
			Currency:       "USD",
		},
		{
			Id:             "H44567",
			Date:           "2024-12-18",
			Ticker:         "META",
			POLT:           0.38,
			Max:            340.75,
			Min:            325.30,
			AvgPrice:       332.02,
			RevenuePercent: 13.0,
			Amount:         85000,
			RevenueBEST:    1100000,
			RevenueTotal:   1250000,
			Currency:       "USD",
		},
		{
			Id:             "H55678",
			Date:           "2024-12-17",
			Ticker:         "NVDA",
			POLT:           0.45,
			Max:            515.25,
			Min:            500.30,
			AvgPrice:       507.78,
			RevenuePercent: 14.2,
			Amount:         90000,
			RevenueBEST:    1400000,
			RevenueTotal:   1600000,
			Currency:       "USD",
		},
		{
			Id:             "H99888",
			Date:           "2024-12-16",
			Ticker:         "NFLX",
			POLT:           0.22,
			Max:            590.10,
			Min:            570.80,
			AvgPrice:       580.00,
			RevenuePercent: 16.3,
			Amount:         65000,
			RevenueBEST:    950000,
			RevenueTotal:   1050000,
			Currency:       "USD",
		},
		{
			Id:             "H22334",
			Date:           "2024-12-15",
			Ticker:         "DIS",
			POLT:           0.30,
			Max:            195.50,
			Min:            180.20,
			AvgPrice:       187.85,
			RevenuePercent: 8.9,
			Amount:         130000,
			RevenueBEST:    1200000,
			RevenueTotal:   1400000,
			Currency:       "USD",
		},
		{
			Id:             "H66789",
			Date:           "2024-12-14",
			Ticker:         "SPY",
			POLT:           0.50,
			Max:            450.00,
			Min:            430.25,
			AvgPrice:       440.12,
			RevenuePercent: 10.0,
			Amount:         110000,
			RevenueBEST:    1300000,
			RevenueTotal:   1450000,
			Currency:       "USD",
		},
	}
	indicators.CalculateIndicators(recordsDemo)
	var recordsFormatted []hrecord.RecordDisplay
	for _, r := range records {
		rd, err := rparser.CreateCurrencyConverted(r)
		if err != nil {
			log.Println("error getting landing page: ", err)
			return c.Render("views/404", nil)
		}
		recordsFormatted = append(recordsFormatted, rd)
	}
	return c.Render("views/home", fiber.Map{
		"Records": recordsFormatted,
	})
}

func CompanyDetails(c *fiber.Ctx) error {
	tkr := strings.ToUpper(c.Params("tkr"))
	company, err := company.GetDetailsByTkr(tkr)
	if err != nil {
		log.Println("error getting landing page: ", err)
		return c.SendStatus(502)
	}
	// records, err := hrecord.GetRecordsByTkrAndTimeframe(tkr, 32)
	// if err != nil {
	// 	log.Println("error getting landing page: ", err)
	// 	return c.SendStatus(502)
	// }
	// indicators.CalculateIndicators(records)
	return c.Render("views/company", fiber.Map{
		"Company": company,
	})
}

func AllCompanies(c *fiber.Ctx) error {
	companies, err := company.GetAll()
	if err != nil {
		log.Println("error getting all companies:", err)
		return c.Render("views/404", nil)
	}
	return c.Render("views/companies_list", fiber.Map{
		"Companies": companies,
	})
}

// Handler to send trading signals over SSE
func CompanyTradingSignals(c *fiber.Ctx) error {
	// Set headers for SSE
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	// Log when the SSE connection is established
	log.Println("SSE connection established")

	// Loop to send trading signals
	for {
		// Simulate trading signal (replace with actual logic)
		tradingSignal := fmt.Sprintf("Trading Signal: %s", time.Now().Format(time.RFC3339))

		// Log the signal being sent
		log.Println("Sending signal:", tradingSignal)

		// Send the trading signal to the client
		err := c.SendString(fmt.Sprintf("data: %s\n\n", tradingSignal))
		if err != nil {
			log.Println("Error sending SSE message:", err)
			break
		}

		// Wait for a few seconds before sending the next signal
		time.Sleep(5 * time.Second)
	}

	return nil
}
