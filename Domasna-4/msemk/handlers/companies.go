package handlers

import (
	"log"
	"msemk/pkg/company"
	"msemk/pkg/queue/dbq"

	"github.com/gofiber/fiber/v2"
)

func CreateCompany(c *fiber.Ctx) error {
	company, err := company.NewCompany(company.WithPersistence(&company.SQLPersistence{}))
	if err != nil {
		log.Printf("error creating a new record: %v", err)
		return c.SendStatus(500)
	}
	err = company.Bind(c.Body())
	if err != nil {
		log.Printf("error binding to record: %v", err)
		return c.Status(400).SendString("Malformed JSON")
	}
	err = company.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dbq.Q.Enqueue(company)
	return c.Status(200).JSON(company)
}

func GetCompany(c *fiber.Ctx) error {
	ticker := c.Params("tkr")
	cmp, err := company.GetDetailsByTkr(ticker)
	if err != nil {
		log.Printf("error getting company by ticker: %v", err)
		return c.SendStatus(500)
	}
	// request day period calculation
	// request week period calculation
	// request month period calculation
	// request news standing
	newCmp := company.NewCompanyDetaildResponse(cmp, dayPeriod string, weekPeriod string, monthPeriod string, newsStanding string)
	return c.Status(200).JSON(newCmp)
}

func GetManyCompanies(c *fiber.Ctx) error {
	companies, err := company.GetAll()
	if err != nil {
		log.Printf("error getting all companies: %v", err)
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(companies)
}

func GetCompaniesTopTen(c *fiber.Ctx) error {
	return nil
}
