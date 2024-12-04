package views

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/company"
	"stocktrust/pkg/hrecord"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LandingPage(c *fiber.Ctx) error {
	records, err := hrecord.GetTopTen()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return c.Render("views/404", nil)
	}
	return c.Render("views/home", fiber.Map{
		"Records": records,
	})
}

func CompanyDetails(c *fiber.Ctx) error {
	tkr := strings.ToUpper(c.Params("tkr"))
	company, err := company.GetDetailsByTkr(tkr)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return c.Render("views/404", nil)
	}
	return c.Render("views/company", fiber.Map{
		"Company": company,
	})
}

func AllCompanies(c *fiber.Ctx) error {
	companies, err := company.GetAll()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return c.Render("views/404", nil)
	}
	return c.Render("views/companies_list", fiber.Map{
		"Companies": companies,
	})
}
