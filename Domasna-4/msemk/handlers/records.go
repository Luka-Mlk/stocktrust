package handlers

import (
	"log"
	"msemk/pkg/hrecord"
	"msemk/pkg/queue/dbq"

	"github.com/gofiber/fiber/v2"
)

func CreateRecord(c *fiber.Ctx) error {
	record, err := hrecord.NewHRecord(hrecord.WithPersistence(&hrecord.SQLPersistence{}))
	if err != nil {
		log.Printf("error creating a new record: %v", err)
		return c.SendStatus(500)
	}
	err = record.Bind(c.Body())
	if err != nil {
		log.Printf("error binding to record: %v", err)
		return c.Status(400).SendString("Malformed JSON")
	}
	err = record.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	dbq.Q.Enqueue(record)
	return c.Status(200).JSON(record)
}

func GetRecordByTicker(c *fiber.Ctx) error {
	// var record hrecord.RecordProxy
	return nil
}

func GetManyRecords(c *fiber.Ctx) error {
	return nil
}

func GetRecordsTopTen(c *fiber.Ctx) error {
	records, err := hrecord.GetTopTen()
	if err != nil {
		log.Printf("error getting top 10 records: %v", err)
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(records)
}
