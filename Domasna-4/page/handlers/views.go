package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func LandingPage(c *fiber.Ctx) error {
	// records, err := hrecord.GetTopTen()
	// if err != nil {
	// 	log.Println("error getting landing page: ", err)
	// 	return c.Render("views/404", nil)
	// }
	// var recordsFormatted []hrecord.RecordDisplay
	// for _, r := range records {
	// 	rd, err := rparser.CreateCurrencyConverted(r)
	// 	if err != nil {
	// 		log.Println("error getting landing page: ", err)
	// 		return c.Render("views/404", nil)
	// 	}
	// 	recordsFormatted = append(recordsFormatted, rd)
	// }
	// return c.Render("views/home", fiber.Map{
	// 	"Records": recordsFormatted,
	// })
}

func ManyCompanies(c *fiber.Ctx) {
	// companies, err := company.GetAll()
	// if err != nil {
	// 	log.Println("error getting all companies:", err)
	// 	return c.Render("views/404", nil)
	// }
	// return c.Render("views/companies_list", fiber.Map{
	// 	"Companies": companies,
	// })
}

func Company(c *fiber.Ctx) error {
	// tkr := strings.ToUpper(c.Params("tkr"))
	// company, err := company.GetDetailsByTkr(tkr)
	// if err != nil {
	// 	log.Println("error getting company details page: ", err)
	// 	return c.Render("views/404", nil)
	// }
	// recordsDayBack, err := hrecord.GetRecordsByTkrAndTimeframe(tkr, 1)
	// if err != nil && err.Error() != "record for ticker not found" {
	// 	log.Println("error getting comapny details page: ", err)
	// 	return c.SendStatus(500)
	// }
	// recordsWeekBack, err := hrecord.GetRecordsByTkrAndTimeframe(tkr, 7)
	// if err != nil && err.Error() != "record for ticker not found" {
	// 	log.Println("error getting comapny details page: ", err)
	// 	return c.SendStatus(500)
	// }
	// recordsMonthBack, err := hrecord.GetRecordsByTkrAndTimeframe(tkr, 32)
	// if err != nil && err.Error() != "record for ticker not found" {
	// 	log.Println("error getting comapny details page: ", err)
	// 	return c.SendStatus(500)
	// }
	// var strategyDay *indicators.Recommendation
	// var strategyWeek *indicators.Recommendation
	// var strategyMonth *indicators.Recommendation
	// if len(recordsDayBack) > 1 {
	// 	strategyDay = indicators.CalculateOscillators(recordsDayBack)
	// }
	// if len(recordsWeekBack) > 1 {
	// 	strategyWeek = indicators.CalculateOscillators(recordsWeekBack)
	// }
	// if len(recordsMonthBack) > 1 {
	// 	strategyMonth = indicators.CalculateOscillators(recordsMonthBack)
	// }
	// return c.Render("views/company", fiber.Map{
	// 	"Company":        company,
	// 	"DayRecommend":   strategyDay,
	// 	"WeekRecommend":  strategyWeek,
	// 	"MonthRecommend": strategyMonth,
	// })
}

func AboutUs(c *fiber.Ctx) error {
	//
}
