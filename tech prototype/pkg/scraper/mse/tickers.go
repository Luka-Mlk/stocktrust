package scraper

import (
	"log"
	"stocktrust/pkg/strings/checker"

	"github.com/gocolly/colly"
)

const (
	dropdown string = ".form-control.dropdown"
	option   string = "option"
	visitUrl string = "https://www.mse.mk/mk/stats/symbolhistory/ADIN"
)

func GetTickers() ([]string, error) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob: "*mse.mk*",
	})
	var tkrs []string
	var cerr error
	c.OnError(func(r *colly.Response, err error) {
		log.Println(err)
		cerr = err
	})
	if cerr != nil {
		log.Println(cerr)
		return nil, cerr
	}
	c.OnHTML(dropdown, func(h *colly.HTMLElement) {
		h.ForEach(option, func(i int, h2 *colly.HTMLElement) {
			tkr := h2.Attr("value")
			if !checker.HasInt(tkr) {
				tkrs = append(tkrs, h2.Attr("value"))
			}
		})
	})
	c.Visit(visitUrl)
	return tkrs, nil
}
