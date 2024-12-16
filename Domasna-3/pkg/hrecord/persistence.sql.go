package hrecord

import (
	"context"
	"errors"
	"fmt"
	"stocktrust/pkg/db"
	"time"
)

type SQLPersistence struct{}

func (p *SQLPersistence) Save(r HRecord) error {
	err := Create(r)
	if err != nil {
		e := fmt.Errorf("error creating HRecord:\n%s", err)
		return e
	}
	return nil
}

func Create(r HRecord) error {
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return e
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
		e := fmt.Errorf("error executing query:\n%s", err)
		return e
	}
	return nil
}

func GetTopTen() ([]HRecord, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return nil, e
	}
	defer db.Release()
	rows, err := db.Query(ctx, getTopTen)
	if err != nil {
		e := fmt.Errorf("error executing query:\n%s", err)
		return nil, e
	}
	var hrecords []HRecord
	for rows.Next() {
		var h HRecord
		var proxy RecordProxy
		err := rows.Scan(
			&proxy.Id,
			&proxy.Date,
			&proxy.Ticker,
			&proxy.POLT,
			&proxy.Max,
			&proxy.Min,
			&proxy.AvgPrice,
			&proxy.RevenuePercent,
			&proxy.Amount,
			&proxy.RevenueBEST,
			&proxy.RevenueTotal,
			&proxy.Currency,
		)
		h.BindFromDB(proxy)
		if err != nil {
			e := fmt.Errorf("error scanning from database:\n%s", err)
			return nil, e
		}
		hrecords = append(hrecords, h)
	}
	if err := rows.Err(); err != nil {
		e := fmt.Errorf("error with rows:\n%s", err)
		return nil, e
	}
	return hrecords, nil
}

func GetLatestTkrDate(tkr string) (time.Time, error) {
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return time.Time{}, e
	}
	defer db.Release()
	rows, err := db.Query(context.Background(), getLatestTickerDate, tkr)
	if err != nil {
		e := fmt.Errorf("error executing query:\n%s", err)
		return time.Time{}, e
	}
	if !rows.Next() {
		return time.Time{}, errors.New("record for ticker not found")
	}
	var data time.Time
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			e := fmt.Errorf("error executing query:\n%s", err)
			return time.Time{}, e
		}
	}
	return data, nil
}

func GetRecordsByTkrAndTimeframe(tkr string, days int) ([]HRecord, error) {
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return nil, e
	}
	defer db.Release()
	now := time.Now()
	recordsFrom := now.AddDate(0, 0, -days)
	recordsFromFormatted := recordsFrom.Format("2006-01-02")
	rows, err := db.Query(
		context.Background(),
		getTimeframeTicker,
		tkr,
		recordsFromFormatted,
	)
	if !rows.Next() {
		return nil, errors.New("record for ticker not found")
	}
	var hrecords []HRecord
	for rows.Next() {
		var h HRecord
		var proxy RecordProxy
		err := rows.Scan(
			&proxy.Id,
			&proxy.Date,
			&proxy.Ticker,
			&proxy.POLT,
			&proxy.Max,
			&proxy.Min,
			&proxy.AvgPrice,
			&proxy.RevenuePercent,
			&proxy.Amount,
			&proxy.RevenueBEST,
			&proxy.RevenueTotal,
		)
		h.BindFromDB(proxy)
		if err != nil {
			e := fmt.Errorf("error scanning from database:\n%s", err)
			return nil, e
		}
		hrecords = append(hrecords, h)
	}
	return hrecords, nil
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

const getTopTen string = `
	SELECT
	    ANY_VALUE(id) AS id,
	    ANY_VALUE(date) AS date,
	    ticker,
	    ANY_VALUE(price_last_transaction) AS price_last_transaction,
	    ANY_VALUE(max) AS max,
	    ANY_VALUE(min) AS min,
	    ANY_VALUE(average_price) AS average_price,
	    ANY_VALUE(revenue_percent) AS revenue_percent,
	    ANY_VALUE(amount) AS amount,
	    ANY_VALUE(revenue_best) AS revenue_best,
	    ANY_VALUE(revenue_total) AS revenue_total,
		ANY_VALUE(currency) AS currency
	FROM history_records
	GROUP BY ticker
	ORDER BY ticker, date DESC
	LIMIT 10;
`

const getLatestTickerDate string = `
	SELECT date
	FROM history_records
	WHERE ticker = $1
	ORDER BY (date) DESC
	LIMIT 1
`

const getTimeframeTicker string = `
SELECT
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
	revenue_total
FROM history_records
WHERE
	ticker = $1
	AND date >= $2;
`
