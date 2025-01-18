package rparser

import (
	"pages/pkg/hrecord"
	hrecfmt "pages/pkg/strings/formatter/hrecord"
)

func CreateCurrencyConverted(hr hrecord.HRecord) (hrecord.RecordDisplay, error) {
	var rd hrecord.RecordDisplay
	rd.Id = hr.Id
	rd.Date = hr.Date
	rd.Ticker = hr.Ticker
	rd.POLT = hrecfmt.FloatToStr(hr.POLT)
	rd.Max = hrecfmt.FloatToStr(hr.Max)
	rd.Min = hrecfmt.FloatToStr(hr.Min)
	rd.AvgPrice = hrecfmt.FloatToStr(hr.AvgPrice)
	rd.RevenuePercent = hrecfmt.FloatToStr(hr.RevenuePercent)
	rd.Amount = hrecfmt.FloatToStr(hr.Amount)
	rd.RevenueBEST = hrecfmt.FloatToStr(hr.RevenueBEST)
	rd.RevenueTotal = hrecfmt.FloatToStr(hr.RevenueTotal)
	rd.Currency = hr.Currency
	return rd, nil
}
