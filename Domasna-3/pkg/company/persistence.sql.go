package company

import (
	"context"
	"fmt"
	"log"
	"stocktrust/pkg/db"
)

type SQLPersistence struct{}

func (p *SQLPersistence) Save(c Company) error {
	tkr, err := GetByTkr(c)
	if err != nil {
		e := fmt.Errorf("error getting ticker for creating compny:\n%s", err)
		return e
	}
	if tkr == c.Ticker {
		return nil
	}
	err = Create(c)
	if err != nil {
		e := fmt.Errorf("error creating company:\n%s", err)
		return e
	}
	return nil
}

func GetByTkr(c Company) (string, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return "", e
	}
	defer db.Release()
	row := db.QueryRow(ctx, getByTicker, c.Ticker)
	db.Release()
	var data string
	row.Scan(&data)
	return data, nil
}

func GetTopCompanies() ([]Company, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return nil, e
	}
	defer db.Release()
	rows, err := db.Query(ctx, getTopCompanies)
	if err != nil {
		e := fmt.Errorf("error executing query:\n%s", err)
		return nil, e
	}
	defer rows.Close()
	var companies []Company
	for rows.Next() {
		var c Company
		err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Address,
			&c.City,
			&c.Country,
			&c.Email,
			&c.Website,
			&c.ContactName,
			&c.ContactPhone,
			&c.ContactEmail,
			&c.Phone,
			&c.Fax,
			&c.Prospect,
			&c.Ticker,
			&c.URL,
		)
		if err != nil {
			e := fmt.Errorf("error scanning from database:\n%s", err)
			return nil, e
		}
		companies = append(companies, c)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}
	return companies, nil
}

func GetDetailsByTkr(tkr string) (*Company, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return nil, e
	}
	defer db.Release()
	row := db.QueryRow(ctx, getDetailsByTicker, tkr)
	if err != nil {
		e := fmt.Errorf("error executing query:\n%s", err)
		return nil, e
	}
	var c Company
	err = row.Scan(
		&c.Id,
		&c.Name,
		&c.Address,
		&c.City,
		&c.Country,
		&c.Email,
		&c.Website,
		&c.ContactName,
		&c.ContactPhone,
		&c.ContactEmail,
		&c.Phone,
		&c.Fax,
		&c.Prospect,
		&c.Ticker,
		&c.URL,
	)
	if err != nil {
		e := fmt.Errorf("error scanning from database:\n%s", err)
		return nil, e
	}
	return &c, nil
}

func GetAll() ([]Company, error) {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return nil, e
	}
	defer db.Release()
	rows, err := db.Query(ctx, getAll)
	if err != nil {
		e := fmt.Errorf("error executing query:\n%s", err)
		return nil, e
	}
	defer rows.Close()
	var companies []Company
	for rows.Next() {
		var c Company
		err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Address,
			&c.City,
			&c.Country,
			&c.Email,
			&c.Website,
			&c.ContactName,
			&c.ContactPhone,
			&c.ContactEmail,
			&c.Phone,
			&c.Fax,
			&c.Prospect,
			&c.Ticker,
			&c.URL,
		)
		if err != nil {
			e := fmt.Errorf("error scanning from database:\n%s", err)
			return nil, e
		}
		companies = append(companies, c)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}
	return companies, nil
}

func Create(c Company) error {
	ctx := context.Background()
	db, err := db.Conn()
	if err != nil {
		e := fmt.Errorf("error connecting to database:\n%s", err)
		return e
	}
	defer db.Release()
	_, err = db.Exec(
		ctx,
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
		e := fmt.Errorf("error executing query:\n%s", err)
		return e
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

const getAll string = `
	SELECT
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
	FROM companies
`

const getTopCompanies string = `
	SELECT
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
	FROM companies
	LIMIT 15
`

const getDetailsByTicker string = `
	SELECT
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
	FROM companies
	WHERE ticker = $1
`
