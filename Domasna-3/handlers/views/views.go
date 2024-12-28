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
		log.Println("error getting company details page: ", err)
		return c.SendStatus(502)
	}
	records, err := hrecord.GetRecordsByTkrAndTimeframe(tkr, 32)
	if err != nil {
		if err.Error() == "record for ticker not found" {
			c.Render("views/404", nil)
		}
		log.Println("error getting comapny details page: ", err)
		return c.SendStatus(502)
	}
	// Calculates indicators - sma, ema, wma, vwma, hma
	indicators.CalculateOscillatorsDay(records)

	return c.Render("views/company", fiber.Map{
		"Company": company,
		// "Recommendation": data,
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
