package company

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/db"
)

type SQLPersistence struct{}

func (p *SQLPersistence) Save(c Company) error {
	tkr, err := GetByTkr(c)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	if tkr == c.Ticker {
		return nil
	}
	err = Create(c)
	if err != nil {
		debug.PrintStack()
		log.Println(err)
		return err
	}
	return nil
}

func GetByTkr(c Company) (string, error) {
	db, err := db.Conn()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return "", err
	}
	row := db.QueryRow(getByTicker, c.Ticker)
	if row == nil {
		return "", nil
	}
	var data string
	row.Scan(&data)
	return data, nil
}

func Create(c Company) error {
	db, err := db.Conn()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	_, err = db.Exec(
		insert,
		c.Id,
		c.Name,
		c.Address,
		c.City,
		c.Country,
		c.Email,
		c.Website,
		c.ContactName,
		c.ContactPhone,
		c.ContactEmail,
		c.Phone,
		c.Fax,
		c.Prospect,
		c.Ticker,
		c.URL,
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
	INSERT INTO companies
		(
			id,
			name,
			address,
			city,
			country,
			email,
			website,
			contact_name,
			contact_phone,
			contact_email,
			phone,
			fax,
			prospect,
			ticker,
			url
		)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
`

const getByTicker string = `
	SELECT ticker
	FROM companies
	WHERE ticker = $1
`
