package scraper

import (
	"errors"
	"fmt"
	"log"
	"stocktrust/pkg/company"
	"stocktrust/pkg/queue/dbq"
	compfmt "stocktrust/pkg/strings/formatter/company"

	"github.com/gocolly/colly"
	"github.com/k0kubun/pp"
)

const (
	accordionSeletor   string = "#collapseFive a:nth-child(1)"
	companyDataWrapper string = "#izdavach"
	noDataCompany      string = "#titleKonf2011"
)

func getCompanyFromTicker(tkr string) error {
	pp.Println(tkr)
	queue := dbq.DBQueue()
	lookupURL := fmt.Sprintf("https://www.mse.mk/mk/search/%s", tkr)
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob: "*mse.mk*",
	})
	cc := c.Clone()
	var cerr error
	c.OnError(func(r *colly.Response, err error) {
		log.Println(err)
		cerr = err
	})
	if cerr != nil {
		log.Println(cerr)
		return cerr
	}
	c.OnHTML(accordionSeletor, func(h *colly.HTMLElement) {
		companyUrl := h.Attr("href")
		cc.Visit("https://mse.mk" + companyUrl)
	})
	cc.OnError(func(r *colly.Response, err error) {
		log.Println(err)
		cerr = err
	})
	if cerr != nil {
		log.Println(cerr)
		return cerr
	}
	var (
		name         string
		address      string
		city         string
		country      string
		email        string
		website      string
		contactName  string
		contactPhone string
		contactEmail string
		phone        string
		fax          string
		prospect     string
		url          string
	)
	cc.OnHTML(companyDataWrapper, func(h *colly.HTMLElement) {
		name = h.ChildText(".title")
		if name == "" {
			return
		}
		address = h.ChildText(companyDataWrapper + " > .row:nth-child(3) .col-md-8")
		if address == "" {
			return
		}
		city = h.ChildText(".row:nth-child(4) .col-md-8")
		country = h.ChildText(".row:nth-child(5) .col-md-8")
		email = h.ChildText(".row:nth-child(6) .col-md-8")
		website = h.ChildText(".row:nth-child(7) .col-md-8")
		contactName = h.ChildText("#popover-content .row:nth-child(1) .col-md-8")
		contactPhone = h.ChildText("#popover-content .row:nth-child(2) .col-md-8")
		contactEmail = h.ChildText("#popover-content .row:nth-child(3) .col-md-8")
		phone = h.ChildText(".row:nth-child(9) .col-md-8")
		fax = h.ChildText(".row:nth-child(10) .col-md-8")
		prospect = h.ChildAttr(".row:nth-child(11) .col-md-8 a", "href")
		url = h.Request.URL.String()
	})
	if cerr != nil {
		return cerr
	}
	cc.OnHTML(noDataCompany, func(h *colly.HTMLElement) {
		name = h.Text
		url = h.Request.URL.String()
	})
	c.Visit(lookupURL)
	cmp, err := company.NewCompany(
		company.WithTicker(tkr),
		company.WithPersistence(&company.SQLPersistence{}),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	cmp.Name = name
	cmp.Address = address
	cmp.City = city
	cmp.Country = country
	cmp.Email = email
	cmp.Website = website
	cmp.ContactName = contactName
	cmp.ContactPhone = contactPhone
	cmp.ContactEmail = contactEmail
	cmp.Phone = phone
	cmp.Fax = fax
	cmp.Prospect = prospect
	cmp.URL = url
	if name == "" || address == "" {
		return errors.New("invalid company")
	}
	compfmt.Company(cmp)
	pp.Println(cmp)

	queue.Enqueue(cmp)

	return nil
}
