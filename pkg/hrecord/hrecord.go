package hrecord

import "log"

type HRecord struct {
	Date           string
	Ticker         string
	POLT           float32
	Max            float32
	Min            float32
	AvgPrice       float32
	RevenuePercent float32
	Amount         int
	RevenueBEST    float32
	RevenueTotal   float32
	Currency       string

	persistences []Persistence
}

type Persistence interface {
	Save(h HRecord) error
}

type Option func(*HRecord) error

func NewHRecord(o ...Option) (*HRecord, error) {
	r := &HRecord{}
	for _, option := range o {
		if err := option(r); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return r, nil
}

func WithDate(date string) Option {
	return func(h *HRecord) error {
		h.Date = date
		return nil
	}
}

func WithTicker(tkr string) Option {
	return func(h *HRecord) error {
		h.Ticker = tkr
		return nil
	}
}

func WithPOLT(price float32) Option {
	return func(h *HRecord) error {
		h.POLT = price
		return nil
	}
}

func WithMax(price float32) Option {
	return func(h *HRecord) error {
		h.Max = price
		return nil
	}
}

func WithMin(price float32) Option {
	return func(h *HRecord) error {
		h.Min = price
		return nil
	}
}

func WithAvgPrice(price float32) Option {
	return func(h *HRecord) error {
		h.AvgPrice = price
		return nil
	}
}

func WithRevenuePercent(price float32) Option {
	return func(h *HRecord) error {
		h.RevenuePercent = price
		return nil
	}
}

func WithAmount(amount int) Option {
	return func(h *HRecord) error {
		h.Amount = amount
		return nil
	}
}

func WithRevenueBEST(price float32) Option {
	return func(h *HRecord) error {
		h.RevenueBEST = price
		return nil
	}
}

func WithRevenueTotal(price float32) Option {
	return func(h *HRecord) error {
		h.RevenueTotal = price
		return nil
	}
}

func WithCurrency(curr string) Option {
	return func(h *HRecord) error {
		h.Currency = curr
		return nil
	}
}

func WithPersistence(ps Persistence) Option {
	return func(h *HRecord) error {
		h.persistences = append(h.persistences, ps)
		return nil
	}
}

func (r *HRecord) Save() error {
	for _, persistence := range r.persistences {
		err := persistence.Save(*r)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
