package hrecord

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"stocktrust/pkg/db"
	"time"
)

type SQLPersistence struct{}

func (p *SQLPersistence) Save(r HRecord) error {
	err := Create(r)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
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
	defer db.Release()
	_, err = db.Exec(
		context.Background(),
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

func GetLatestTkrDate(tkr string) (time.Time, error) {
	db, err := db.Conn()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return time.Time{}, err
	}
	defer db.Release()
	rows, err := db.Query(context.Background(), getLatestTickerDate, tkr)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return time.Time{}, err
	}
	if !rows.Next() {
		return time.Time{}, errors.New("record for ticker not found")
	}
	var data time.Time
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Println(err)
		}
	}
	return data, nil
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

const getLatestTickerDate string = `
	SELECT date
	FROM history_records
	WHERE ticker = $1
	ORDER BY (date) DESC
	LIMIT 1
`