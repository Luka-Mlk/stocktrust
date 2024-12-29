package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/queue/dbq"
	hrecfmt "stocktrust/pkg/strings/formatter/hrecord"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func updateHrForTicker(tkr string, lDate time.Time) error {
	end := fmt.Sprintf("https://www.mse.mk/mk/stats/symbolhistory/%s", tkr)
	ctyp := "application/x-www-form-urlencoded"
	cdate := time.Now()
	data := url.Values{}
	datesMatch := false
	for !datesMatch {
		lDatePlus := lDate.AddDate(0, 0, 365)
		data.Set("FromDate", lDate.Format("02.01.2006"))
		data.Set("ToDate", lDatePlus.Format("02.01.2006"))
		data.Set("Code", tkr)
		res, err := http.Post(end, ctyp, strings.NewReader(data.Encode()))
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		fName := "pkg/scraper/mse/html/history.html"
		err = os.WriteFile(fName, body, 0660)
		if err != nil {
			log.Print(err)
			debug.PrintStack()
			return err
		}
		err = scrapeFile(fName, tkr)
		if err != nil {
			log.Println(err)
			return err
		}
		lDate = lDatePlus
		if lDatePlus.After(cdate) {
			datesMatch = true
		}
	}
	return nil
}

func getHrListForTicker(tkr string, group int) error {
	end := fmt.Sprintf("https://www.mse.mk/mk/stats/symbolhistory/%s", tkr)
	ctyp := "application/x-www-form-urlencoded"
	cdate := time.Now()
	data := url.Values{}
	for i := 0; i < 10; i++ {
		cdateOneLess := cdate.AddDate(0, 0, -365)
		data.Set("FromDate", cdateOneLess.Format("02.01.2006"))
		data.Set("ToDate", cdate.Format("02.01.2006"))
		data.Set("Code", tkr)
		res, err := http.Post(end, ctyp, strings.NewReader(data.Encode()))
		if err != nil {
			log.Println(err)
			return err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		fName := fmt.Sprintf("pkg/scraper/mse/html/history-group-%v.html", group)
		err = os.WriteFile(fName, body, 0660)
		if err != nil {
			log.Print(err)
			debug.PrintStack()
			return err
		}
		err = scrapeFile(fName, tkr)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		cdate = cdate.AddDate(0, 0, -365)
	}
	return nil
}

func scrapeFile(file string, tkr string) error {
	queue := dbq.DBQueue()
	c := colly.NewCollector()
	c.WithTransport(http.NewFileTransport(http.Dir("./")))
	c.Limit(&colly.LimitRule{
		DomainGlob: "*mse.mk*",
	})
	// collector error
	var cerr error
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			cerr = err
		}
	})
	if cerr != nil {
		return cerr
	}
	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			date := h.ChildText("td:nth-child(1)")
			date = hrecfmt.FormatDate(date)
			polt := h.ChildText("td:nth-child(2)")
			poltFloat, err := hrecfmt.EUDecimalToUSFromStr(polt)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			max := h.ChildText("td:nth-child(3)")
			maxFloat, err := hrecfmt.EUDecimalToUSFromStr(max)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			min := h.ChildText("td:nth-child(4)")
			minFloat, err := hrecfmt.EUDecimalToUSFromStr(min)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			avgprice := h.ChildText("td:nth-child(5)")
			avgPriceFloat, err := hrecfmt.EUDecimalToUSFromStr(avgprice)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			revenuePerc := h.ChildText("td:nth-child(6)")
			revenuePercFloat, err := hrecfmt.EUDecimalToUSFromStr(revenuePerc)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			amount := h.ChildText("td:nth-child(7)")
			amountFloat, err := hrecfmt.EUDecimalToUSFromStr(amount)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			// UNCOMMENT HERE TO SCRAPE FASTER
			// LEAVE COMMENTED TO SCRAPE MORE
			if amountFloat == 0 {
				return
			}
			revBest := h.ChildText("td:nth-child(8)")
			revBestFloat, err := hrecfmt.EUDecimalToUSFromStr(revBest)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			revTot := h.ChildText("td:nth-child(9)")
			revTotFloat, err := hrecfmt.EUDecimalToUSFromStr(revTot)
			if err != nil {
				log.Println(err)
				cerr = err
				return
			}
			hr, err := hrecord.NewHRecord(
				hrecord.WithDate(date),
				hrecord.WithTicker(tkr),
				hrecord.WithPOLT(float32(poltFloat)),
				hrecord.WithMax(float32(maxFloat)),
				hrecord.WithMin(float32(minFloat)),
				hrecord.WithAvgPrice(float32(avgPriceFloat)),
				hrecord.WithRevenuePercent(float32(revenuePercFloat)),
				hrecord.WithAmount(amountFloat),
				hrecord.WithRevenueBEST(float32(revBestFloat)),
				hrecord.WithRevenueTotal(float32(revTotFloat)),
				hrecord.WithCurrency("MKD"),
				hrecord.WithPersistence(&hrecord.SQLPersistence{}))
			if err != nil {
				cerr = err
			}
			queue.Enqueue(hr)
		})
	})
	if cerr != nil {
		return cerr
	}
	c.Visit(file)
	return nil
}
