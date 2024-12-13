package hrecord

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/rs/xid"
)

type HRecord struct {
	Id             string
	Date           string
	Ticker         string
	POLT           float32
	Max            float32
	Min            float32
	AvgPrice       float32
	RevenuePercent float32
	Amount         float32
	RevenueBEST    float32
	RevenueTotal   float32
	Currency       string

	persistences []Persistence
}

type RecordProxy struct {
	Id             string
	Date           time.Time
	Ticker         string
	POLT           string
	Max            string
	Min            string
	AvgPrice       string
	RevenuePercent string
	Amount         string
	RevenueBEST    string
	RevenueTotal   string
	Currency       string
}

type RecordDisplay struct {
	Id             string
	Date           string
	Ticker         string
	POLT           string
	Max            string
	Min            string
	AvgPrice       string
	RevenuePercent string
	Amount         string
	RevenueBEST    string
	RevenueTotal   string
	Currency       string
}

type Persistence interface {
	Save(h HRecord) error
}

type Option func(*HRecord) error

func NewHRecord(o ...Option) (*HRecord, error) {
	r := &HRecord{
		Id: xid.New().String(),
	}
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

func WithAmount(amount float32) Option {
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

func (r *HRecord) BindFromDB(proxy RecordProxy) error {
	floatPOLT, err := strconv.ParseFloat(proxy.POLT, 32)
	if err != nil {
		e := fmt.Errorf("error parsing POLT:\n%s", err)
		return e
	}
	r.POLT = float32(floatPOLT)
	floatMax, err := strconv.ParseFloat(proxy.Max, 32)
	if err != nil {
		e := fmt.Errorf("error parsing max:\n%s", err)
		return e
	}
	r.Max = float32(floatMax)
	floatMin, err := strconv.ParseFloat(proxy.Min, 32)
	if err != nil {
		e := fmt.Errorf("error parsing min:\n%s", err)
		return e
	}
	r.Min = float32(floatMin)
	floatAvgPrice, err := strconv.ParseFloat(proxy.AvgPrice, 32)
	if err != nil {
		e := fmt.Errorf("error parsing avg:\n%s", err)
		return e
	}
	r.AvgPrice = float32(floatAvgPrice)
	floatRevenuePercent, err := strconv.ParseFloat(proxy.RevenuePercent, 32)
	if err != nil {
		e := fmt.Errorf("error parsing revenue percent:\n%s", err)
		return e
	}
	r.RevenuePercent = float32(floatRevenuePercent)
	floatAmount, err := strconv.ParseFloat(proxy.Amount, 32)
	if err != nil {
		e := fmt.Errorf("error parsing revenue amount:\n%s", err)
		return e
	}
	r.Amount = float32(floatAmount)
	floatRevenueBEST, err := strconv.ParseFloat(proxy.RevenueBEST, 32)
	if err != nil {
		e := fmt.Errorf("error parsing revenue revenue BEST:\n%s", err)
		return e
	}
	r.RevenueBEST = float32(floatRevenueBEST)
	floatRevenueTotal, err := strconv.ParseFloat(proxy.RevenueTotal, 32)
	if err != nil {
		e := fmt.Errorf("error parsing revenue total:\n%s", err)
		return e
	}
	r.RevenueTotal = float32(floatRevenueTotal)

	r.Id = proxy.Id
	r.Date = proxy.Date.Format("2006-01-02")
	r.Ticker = proxy.Ticker
	r.Currency = proxy.Currency
	return nil
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
