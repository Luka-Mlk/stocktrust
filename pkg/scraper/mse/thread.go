package scraper

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/hrecord"
	"sync"
)

func divideLoad(wg *sync.WaitGroup, tkrs []string, group int) error {
	defer wg.Done()
	for _, tkr := range tkrs {
		// FILTER NO 2 - get latest record date
		latestDate, err := hrecord.GetLatestTkrDate(tkr)
		if err != nil && err.Error() == "record for ticker not found" {
			err = getCompanyFromTicker(tkr)
			if err != nil {
				continue
			}
			err = getHrListForTicker(tkr)
			if err != nil {
				log.Println(err)
				debug.PrintStack()
				return err
			}
		} else if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		} else {
			// FILTER NO 3 - get all missing records
			err = updateHrForTicker(tkr, latestDate)
			if err != nil {
				log.Println(err)
				debug.PrintStack()
				return err
			}
		}
	}
	return nil
}
