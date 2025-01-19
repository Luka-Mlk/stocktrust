package handlers

import (
	"log"
	"msemk/pkg/company"
	"msemk/pkg/hrecord"
	"msemk/pkg/indicators"
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
	recordsDay, err := hrecord.GetRecordsByTkrAndTimeframe(ticker, 1)
	if err != nil && err.Error() != "record for ticker not found" {
		log.Printf("error getting records by ticker %s and timeframe %d: %v\n", ticker, 30, err)
		return c.SendStatus(500)
	}
	recordsWeek, err := hrecord.GetRecordsByTkrAndTimeframe(ticker, 7)
	if err != nil && err.Error() != "record for ticker not found" {
		log.Printf("error getting records by ticker %s and timeframe %d: %v\n", ticker, 30, err)
		return c.SendStatus(500)
	}
	recordsMonth, err := hrecord.GetRecordsByTkrAndTimeframe(ticker, 30)
	if err != nil && err.Error() != "record for ticker not found" {
		log.Printf("error getting records by ticker %s and timeframe %d: %v\n", ticker, 30, err)
		return c.SendStatus(500)
	}
	var strategyDay indicators.Recommendation
	var strategyWeek indicators.Recommendation
	var strategyMonth indicators.Recommendation
	if len(recordsDay) > 1 {
		strategyDay = indicators.CalculateOscillators(recordsDay)
	}
	if len(recordsWeek) > 1 {
		strategyWeek = indicators.CalculateOscillators(recordsWeek)
	}
	if len(recordsMonth) > 1 {
		strategyMonth = indicators.CalculateOscillators(recordsMonth)
	}
	newCmp := company.NewCompanyDetaildResponse(cmp, strategyDay, strategyWeek, strategyMonth, "newsStanding")
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
