package handlers

import (
	"log"
	"pages/pkg/company"
	"pages/pkg/hrecord"
	rparser "pages/pkg/strings/parser/records"
	"strings"

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

func ManyCompanies(c *fiber.Ctx) error {
	companies, err := company.GetManyCompanies()
	if err != nil {
		log.Println("error getting all companies:", err)
		return c.Render("views/404", nil)
	}
	return c.Render("views/companies_list", fiber.Map{
		"Companies": companies,
	})
}

func Company(c *fiber.Ctx) error {
	ticker := strings.ToUpper(c.Params("tkr"))
	company, err := company.GetCompanyByTicker(ticker)
	if err != nil {
		log.Println("error getting company details page: ", err)
		return c.Render("views/404", nil)
	}
	return c.Render("views/company", fiber.Map{
		"Company": company,
	})
}

func AboutUs(c *fiber.Ctx) error {
	//
	return nil
}
