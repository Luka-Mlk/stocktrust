package scraper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"stocktrust/pkg/hrecord"
	"stocktrust/pkg/queue/dbq"
	hrecfmt "stocktrust/pkg/strings/formatter/hrecord"
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/net/html"
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
		// Send request and get HTML response in memory
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
		// Directly process the HTML content without saving it to a file
		err = scrapeHTMLContent(body, tkr)
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

func getHrListForTicker(tkr string) error {
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
		err = scrapeHTMLContent2(body, tkr)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		cdate = cdate.AddDate(0, 0, -365)
	}
	return nil
}

func scrapeHTMLContent2(content []byte, tkr string) error {
	// queue := dbq.DBQueue()
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return err
	}
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "tbody" {
			pp.Println(node.FirstChild)
		}
	}
	crawler(doc)
	return nil
}

func scrapeHTMLContent(content []byte, tkr string) error {
	queue := dbq.DBQueue()
	// Parse the HTML content using net/html package
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return err
	}
	// Define the traversal function to go through the HTML tree
	var cerr error
	var traverseNode func(*html.Node)
	traverseNode = func(n *html.Node) {
		pp.Println("Start the traversal from the root node (document root)")
		// We want to look specifically for the <tbody> element and its children <tr> elements
		if n.Type == html.ElementNode && n.Data == "tbody" {
			// Now, process the <tr> elements inside <tbody>
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "tr" {
					// Extract the data from the <td> elements within this <tr>
					var date, polt, max, min, avgprice, revenuePerc, amount, revBest, revTot string
					// Traverse child nodes (each <td> element)
					tdIndex := 1 // This is to track which <td> we are on
					for td := c.FirstChild; td != nil; td = td.NextSibling {
						if td.Type == html.ElementNode && td.Data == "td" {
							// Ensure we have text content in the <td>
							if td.FirstChild != nil {
								textContent := strings.TrimSpace(td.FirstChild.Data)
								// Use switch based on the <td> position (1st, 2nd, etc.)
								switch tdIndex {
								case 1:
									date = textContent
								case 2:
									polt = textContent
								case 3:
									max = textContent
								case 4:
									min = textContent
								case 5:
									avgprice = textContent
								case 6:
									revenuePerc = textContent
								case 7:
									amount = textContent
								case 8:
									revBest = textContent
								case 9:
									revTot = textContent
								}
								tdIndex++
							}
						}
					}
					// If all data fields are found, parse and enqueue the data
					if date != "" && polt != "" && max != "" && min != "" && avgprice != "" && revenuePerc != "" && amount != "" && revBest != "" && revTot != "" {
						date = hrecfmt.FormatDate(date)
						// Parse the decimal values from string to float
						poltFloat, err := hrecfmt.EUDecimalToUSFromStr(polt)
						if err != nil {
							log.Println("Error parsing POLT:", err)
							cerr = err
							return
						}
						maxFloat, err := hrecfmt.EUDecimalToUSFromStr(max)
						if err != nil {
							log.Println("Error parsing max:", err)
							cerr = err
							return
						}
						minFloat, err := hrecfmt.EUDecimalToUSFromStr(min)
						if err != nil {
							log.Println("Error parsing min:", err)
							cerr = err
							return
						}
						avgPriceFloat, err := hrecfmt.EUDecimalToUSFromStr(avgprice)
						if err != nil {
							log.Println("Error parsing avgPrice:", err)
							cerr = err
							return
						}
						revenuePercFloat, err := hrecfmt.EUDecimalToUSFromStr(revenuePerc)
						if err != nil {
							log.Println("Error parsing revenuePerc:", err)
							cerr = err
							return
						}
						amountFloat, err := hrecfmt.EUDecimalToUSFromStr(amount)
						if err != nil {
							log.Println("Error parsing amount:", err)
							cerr = err
							return
						}
						revBestFloat, err := hrecfmt.EUDecimalToUSFromStr(revBest)
						if err != nil {
							log.Println("Error parsing revBest:", err)
							cerr = err
							return
						}
						revTotFloat, err := hrecfmt.EUDecimalToUSFromStr(revTot)
						if err != nil {
							log.Println("Error parsing revTot:", err)
							cerr = err
							return
						}
						// Create a new historical record (HRecord)
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
						// Enqueue the HRecord to the DBQueue
						queue.Enqueue(hr)
					}
				}
			}
		}
		// Continue traversing child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			pp.Println("Continue traversing child nodes")
			traverseNode(c)
		}
	}
	// Start the traversal from the root node (document root)
	traverseNode(doc)
	if cerr != nil {
		return cerr
	}
	return nil
}
