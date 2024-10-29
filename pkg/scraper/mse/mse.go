package scraper

import (
	"stocktrust/pkg/company"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/hrecordlist"
)

const (
	link1 string = "https://www.mse.mk/mk/stats/symbolhistory/KMB"
)

func Init() {
	tkrs := getTickers()
	for _, tkr := range tkrs {
		getCompanyFromTicker(tkr)
		getHRForTicker(tkr)
	}
}

func Update() {
	tkrs := getTickers()
	for _, tkr := range tkrs {
		getCompanyFromTicker(tkr)
		getLatestForTicker(tkr)
	}
}

func getTickers() []string {
	return nil
}

func getCompanyFromTicker(tkr string) []company.Company {
	return nil
}

func getHRForTicker(tkr string) []hrecordlist.HRecordList {
	return nil
}

func getLatestForTicker(tkr string) hrecord.HRecord {
	return hrecord.HRecord{}
}
