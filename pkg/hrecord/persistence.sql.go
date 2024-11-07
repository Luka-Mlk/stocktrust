package hrecord

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/db"
)

type SQLPersistence struct{}

func (p *SQLPersistence) Save(r HRecord) error {
	err := Create(r)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Create(r HRecord) error {
	db, err := db.Conn()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.Exec(
		insert,
		r.Id,
		r.Date,
		r.Ticker,
		r.POLT,
		r.Max,
		r.Min,
		r.AvgPrice,
		r.RevenuePercent,
		r.Amount,
		r.RevenueBEST,
		r.RevenueTotal,
		r.Currency,
	)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

// ======== SQL QUERIES ========

const insert string = `
	INSERT INTO history_records 
		(
			id,
			date, 
			ticker, 
			price_last_transaction, 
			max, 
			min, 
			average_price, 
			revenue_percent, 
			amount,
			revenue_best,
			revenue_total,
			currency
		)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
