package views

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/company"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LandingPage(c *fiber.Ctx) error {
	companies, err := company.GetTopCompanies()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return c.Render("views/404", nil)
	}
	return c.Render("views/home", fiber.Map{
		"Companies": companies,
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
