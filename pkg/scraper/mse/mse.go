package scraper

import (
	"log"
	"runtime/debug"
)

func Init() error {
	// dbtkrs := map[string]date
	// if dbtkrs != empty
	// goroutine for ticker history that already exists in the dbtkrs map?

	// else map empty execute code below
	tkrs, err := GetTickers()
	if err != nil {
		log.Println(err)
		return err
	}
	// update function call here? on history
	// var cl []*company.Company
	for _, tkr := range tkrs {
		// if tkrmem[tkr] exist run update from it's value
		// check if current ticker is in the slice
		// if yes check latest date and get data from latest date
		// if not get entire history for ticker going 10 yr back
		_, err := getCompanyFromTicker(tkr)
		if err != nil {
			continue
			// log.Println(err)
			// debug.PrintStack()
			// return err
			// cl = append(cl, company)
		}
		_, err = getHrListForTicker(tkr)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
	}
	return nil
}

func Update() {
	// tkrs := getTickers()
	// for _, tkr := range tkrs {
	// 	getCompanyFromTicker(tkr)

	// 	getLatestForTicker(tkr)
	// }
}

// func getLatestForTicker(tkr string) hrecord.HRecord {
// 	return hrecord.HRecord{}
// }
