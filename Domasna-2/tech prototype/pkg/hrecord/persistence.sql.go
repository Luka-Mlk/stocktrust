package hrecord

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"stocktrust/pkg/db"
	"strconv"
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

func GetTopTen() ([]HRecord, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return nil, err
	}
	defer db.Release()
	rows, err := db.Query(ctx, getTopTen)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return nil, err
	}
	var hrecords []HRecord
	for rows.Next() {
		var h HRecord
		proxy := struct {
			id             string
			date           time.Time
			ticker         string
			POLT           string
			max            string
			min            string
			avgPrice       string
			revenuePercent string
			amount         string
			revenueBEST    string
			revenueTotal   string
			currency       string
		}{}
		err := rows.Scan(
			&proxy.id,
			&proxy.date,
			&proxy.ticker,
			&proxy.POLT,
			&proxy.max,
			&proxy.min,
			&proxy.avgPrice,
			&proxy.revenuePercent,
			&proxy.amount,
			&proxy.revenueBEST,
			&proxy.revenueTotal,
			&proxy.currency,
		)

		floatPOLT, err := strconv.ParseFloat(proxy.POLT, 32)
		if err != nil {
			log.Println("Error parsing POLT:", err)
			debug.PrintStack()
			return nil, err
		}
		h.POLT = float32(floatPOLT)

		// Max
		floatMax, err := strconv.ParseFloat(proxy.max, 32)
		if err != nil {
			log.Println("Error parsing Max:", err)
			debug.PrintStack()
			return nil, err
		}
		h.Max = float32(floatMax)

		// Min
		floatMin, err := strconv.ParseFloat(proxy.min, 32)
		if err != nil {
			log.Println("Error parsing Min:", err)
			debug.PrintStack()
			return nil, err
		}
		h.Min = float32(floatMin)

		// AvgPrice
		floatAvgPrice, err := strconv.ParseFloat(proxy.avgPrice, 32)
		if err != nil {
			log.Println("Error parsing AvgPrice:", err)
			debug.PrintStack()
			return nil, err
		}
		h.AvgPrice = float32(floatAvgPrice)

		// RevenuePercent
		floatRevenuePercent, err := strconv.ParseFloat(proxy.revenuePercent, 32)
		if err != nil {
			log.Println("Error parsing RevenuePercent:", err)
			debug.PrintStack()
			return nil, err
		}
		h.RevenuePercent = float32(floatRevenuePercent)

		// Amount
		floatAmount, err := strconv.ParseFloat(proxy.amount, 32)
		if err != nil {
			log.Println("Error parsing Amount:", err)
			debug.PrintStack()
			return nil, err
		}
		h.Amount = float32(floatAmount)

		// RevenueBEST
		floatRevenueBEST, err := strconv.ParseFloat(proxy.revenueBEST, 32)
		if err != nil {
			log.Println("Error parsing RevenueBEST:", err)
			debug.PrintStack()
			return nil, err
		}
		h.RevenueBEST = float32(floatRevenueBEST)

		// RevenueTotal
		floatRevenueTotal, err := strconv.ParseFloat(proxy.revenueTotal, 32)
		if err != nil {
			log.Println("Error parsing RevenueTotal:", err)
			debug.PrintStack()
			return nil, err
		}
		h.RevenueTotal = float32(floatRevenueTotal)

		h.Id = proxy.id
		h.Date = proxy.date.Format("2006-02-01")
		h.Ticker = proxy.ticker
		h.Currency = proxy.currency
		if err != nil {
			log.Println("Error scanning row:", err)
			debug.PrintStack()
			return nil, err
		}
		// pp.Println(proxy.ticker)
		// pp.Println("floatPOLT:", floatPOLT)
		// pp.Println("floatMax:", floatMax)
		// pp.Println("floatMin:", floatMin)
		// pp.Println("floatAvgPrice:", floatAvgPrice)
		// pp.Println("floatRevenuePercent:", floatRevenuePercent)
		// pp.Println("floatAmount:", floatAmount)
		// pp.Println("floatRevenueBEST:", floatRevenueBEST)
		// pp.Println("floatRevenueTotal:", floatRevenueTotal)
		// return nil, nil
		hrecords = append(hrecords, h)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		debug.PrintStack()
		return nil, err
	}
	return hrecords, nil
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

const getTopTen string = `
	SELECT DISTINCT ON (ticker)
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
	FROM history_records
	WHERE revenue_best > 10000
	ORDER BY ticker, revenue_best DESC
	LIMIT 10;
`

const getLatestTickerDate string = `
	SELECT date
	FROM history_records
	WHERE ticker = $1
	ORDER BY (date) DESC
	LIMIT 1
`
