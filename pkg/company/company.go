package company

import "log"

type Company struct {
	Name     string
	Address  string
	City     string
	Country  string
	Email    string
	Website  string
	Contact  string
	Phone    string
	Fax      string
	Prospect string
	Ticker   string

	persistences []Persistence
}

type Persistence interface {
	Save(c Company) error
}

type Option func(*Company) error

func NewCompany(o ...Option) (*Company, error) {
	c := &Company{}
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

func WithContact(contact string) Option {
	return func(c *Company) error {
		c.Contact = contact
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
