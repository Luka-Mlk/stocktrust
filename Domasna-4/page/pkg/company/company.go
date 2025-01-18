package company

import (
	"log"
	"pages/pkg/indicators"

	"github.com/rs/xid"
)

type Company struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	City         string `json:"city" validate:"required"`
	Country      string `json:"country" validate:"required"`
	Email        string `json:"email" validate:"email"`
	Website      string `json:"website"`
	ContactName  string `json:"contact_name"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email" validate:"email"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Prospect     string `json:"prospect"`
	Ticker       string `json:"ticker" validate:"alphanum"`
	URL          string `json:"url" validate:"required"`

	persistences []Persistence
}

type CompanyDetailedResponse struct {
	Id           string                    `json:"id,omitempty"`
	Name         string                    `json:"name" validate:"required"`
	Address      string                    `json:"address" validate:"required"`
	City         string                    `json:"city" validate:"required"`
	Country      string                    `json:"country" validate:"required"`
	Email        string                    `json:"email" validate:"email"`
	Website      string                    `json:"website"`
	ContactName  string                    `json:"contact_name"`
	ContactPhone string                    `json:"contact_phone"`
	ContactEmail string                    `json:"contact_email" validate:"email"`
	Phone        string                    `json:"phone"`
	Fax          string                    `json:"fax"`
	Prospect     string                    `json:"prospect"`
	Ticker       string                    `json:"ticker" validate:"alphanum"`
	URL          string                    `json:"url" validate:"required"`
	DayPeriod    indicators.Recommendation `json:"day_period_recommendation"`
	WeekPeriod   indicators.Recommendation `json:"week_period_recommendation"`
	MonthPeriod  indicators.Recommendation `json:"month_period_recommendation"`
	NewsStanding string                    `json:"news_standing"`
}

type Persistence interface {
	Save(c Company) error
}

type Option func(*Company) error

func NewCompany(o ...Option) (*Company, error) {
	c := &Company{
		Id: xid.New().String(),
	}
	for _, option := range o {
		err := option(c)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return c, nil
}

func WithName(name string) Option {
	return func(c *Company) error {
		c.Name = name
		return nil
	}
}

func WithAddress(adr string) Option {
	return func(c *Company) error {
		c.Address = adr
		return nil
	}
}

func WithCity(cty string) Option {
	return func(c *Company) error {
		c.City = cty
		return nil
	}
}

func WithCountry(ctry string) Option {
	return func(c *Company) error {
		c.Country = ctry
		return nil
	}
}

func WithEmail(email string) Option {
	return func(c *Company) error {
		c.Email = email
		return nil
	}
}

func WithWebsite(website string) Option {
	return func(c *Company) error {
		c.Website = website
		return nil
	}
}

func WithContactName(name string) Option {
	return func(c *Company) error {
		c.ContactName = name
		return nil
	}
}

func WithContactPhone(phone string) Option {
	return func(c *Company) error {
		c.ContactPhone = phone
		return nil
	}
}

func WithContactEmail(email string) Option {
	return func(c *Company) error {
		c.ContactEmail = email
		return nil
	}
}

func WithPhone(phone string) Option {
	return func(c *Company) error {
		c.Phone = phone
		return nil
	}
}

func WithFax(fax string) Option {
	return func(c *Company) error {
		c.Fax = fax
		return nil
	}
}

func WithProspect(prospect string) Option {
	return func(c *Company) error {
		c.Prospect = prospect
		return nil
	}
}

func WithTicker(ticker string) Option {
	return func(c *Company) error {
		c.Ticker = ticker
		return nil
	}
}

func WithURL(url string) Option {
	return func(c *Company) error {
		c.URL = url
		return nil
	}
}

func WithPersistence(ps Persistence) Option {
	return func(h *Company) error {
		h.persistences = append(h.persistences, ps)
		return nil
	}
}

func (c *Company) Save() error {
	for _, persistence := range c.persistences {
		err := persistence.Save(*c)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
