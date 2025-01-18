package handlers

import (
	scraper "stocktrust/pkg/scraper/mse"

	"github.com/gofiber/fiber/v2"
)

func Scraper(c *fiber.Ctx) error {
	go scraper.Init()
	return c.SendStatus(200)
}
