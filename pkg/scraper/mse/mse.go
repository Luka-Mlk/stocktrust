package scraper

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/hrecord"
)

func Init() error {
	// dbtkrs := map[string]date
	// if dbtkrs != empty
	// goroutine for ticker history that already exists in the dbtkrs map?

	// else map empty execute code below
	tkrs, err := GetTickers()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	// update function call here? on history
	for _, tkr := range tkrs {
		latestDate, err := hrecord.GetLatestTkrDate(tkr)
		if err != nil && err.Error() == "record for ticker not found" {
			// Check if invalid company is exact error
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
