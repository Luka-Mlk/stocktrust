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
		latestDate, err := hrecord.GetLatestTkrDate(tkr)
		if err != nil && err.Error() == "record for ticker not found" {
			err = getCompanyFromTicker(tkr)
			if err != nil {
				continue
			}
			err = getHrListForTicker(tkr, group)
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
