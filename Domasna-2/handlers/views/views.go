package views

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/company"

	"github.com/gofiber/fiber/v2"
	"github.com/k0kubun/pp/v3"
)

// LandingPage handler
func LandingPage(c *fiber.Ctx) error {
	companies, err := company.GetAll()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return c.Render("home", fiber.Map{
		"Companies": companies,
	})
}

func GetByTicker(c *fiber.Ctx) error {
	company, err := company.GetDetailsByTkr(c.Params("tkr"))
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	pp.Println(company)
	return c.Render("single", fiber.Map{
		"Company": company,
	})
}
