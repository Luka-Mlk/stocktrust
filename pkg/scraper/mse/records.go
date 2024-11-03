package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/hrecordlist"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/k0kubun/pp"
)

func getHrListForTicker(tkr string) (*hrecordlist.HRecordList, error) {
	end := fmt.Sprintf("https://www.mse.mk/mk/stats/symbolhistory/%s", tkr)
	ctyp := "application/x-www-form-urlencoded"
	cdate := time.Now()
	data := url.Values{}
	for i := 0; i < 11; i++ {
		cdateOneLess := cdate.AddDate(0, 0, -365)
		data.Set("FromDate", cdateOneLess.Format("02.01.2006"))
		data.Set("ToDate", cdate.Format("02.01.2006"))
		data.Set("Code", tkr)
		res, err := http.Post(end, ctyp, strings.NewReader(data.Encode()))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		fName := "pkg/scraper/mse/html/history.html"
		err = os.WriteFile(fName, body, 0660)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		hrl, err := hrecordlist.NewHRecordList(
			hrecordlist.WithPersistence(&hrecordlist.SQLPersistence{}),
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		err = scrapeFile(fName, tkr, hrl)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cdate = cdate.AddDate(0, 0, -365)
	}
	return nil, nil
}

func scrapeFile(file string, tkr string, hrl *hrecordlist.HRecordList) error {
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
	pp.Println("before OnHTML tbody")
	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		pp.Println("found tbody")
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			pp.Println("found tr in tbody")
			date := h.ChildText("td:nth-child(1)")
			polt := h.ChildText("td:nth-child(2)")
			poltFloat, err := strconv.ParseFloat(polt, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			max := h.ChildText("td:nth-child(3)")
			maxFloat, err := strconv.ParseFloat(max, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			min := h.ChildText("td:nth-child(4)")
			minFloat, err := strconv.ParseFloat(min, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			avgprice := h.ChildText("td:nth-child(4)")
			avgPriceFloat, err := strconv.ParseFloat(avgprice, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			revenuePerc := h.ChildText("td:nth-child(5)")
			revenuePercFloat, err := strconv.ParseFloat(revenuePerc, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			amount := h.ChildText("td:nth-child(6)")
			amountInt, err := strconv.Atoi(amount)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			revBest := h.ChildText("td:nth-child(7)")
			revBestFloat, err := strconv.ParseFloat(revBest, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			revTot := h.ChildText("td:nth-child(8)")
			revTotFloat, err := strconv.ParseFloat(revTot, 32)
			if err != nil {
				log.Println(err)
				cerr = err
			}
			hr, err := hrecord.NewHRecord(
				hrecord.WithDate(date),
				hrecord.WithTicker(tkr),
				hrecord.WithPOLT(float32(poltFloat)),
				hrecord.WithMax(float32(maxFloat)),
				hrecord.WithMin(float32(minFloat)),
				hrecord.WithAvgPrice(float32(avgPriceFloat)),
				hrecord.WithRevenuePercent(float32(revenuePercFloat)),
				hrecord.WithAmount(amountInt),
				hrecord.WithRevenueBEST(float32(revBestFloat)),
				hrecord.WithRevenueTotal(float32(revTotFloat)),
				hrecord.WithCurrency("MKD"),
				hrecord.WithPersistence(&hrecord.SQLPersistence{}))
			if err != nil {
				cerr = err
			}
			pp.Println(hr)
			hrl.Append(*hr)
		})
	})
	if cerr != nil {
		return cerr
	}
	c.Visit(file)
	return nil
}
