package scraper

import (
	"fmt"
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
			err = getHrListForTicker(tkr, group)
			if err != nil {
				e := fmt.Errorf("error getting history for ticker %s:\n%s", tkr, err)
				return e
			}
		} else if err != nil {
			e := fmt.Errorf("error getting latest ticker date %s:\n%s", tkr, err)
			return e
		} else {
			// FILTER NO 3 - get all missing records
			err = updateHrForTicker(tkr, latestDate)
			if err != nil {
				e := fmt.Errorf("error updating history for %s:\n%s", tkr, err)
				return e
			}
		}
	}
	return nil
}
