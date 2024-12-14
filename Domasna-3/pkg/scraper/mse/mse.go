package scraper

import (
	"fmt"
	"os"
	"stocktrust/pkg/queue/dbq"
	"strconv"
	"sync"
)

func Init() error {
	q := dbq.DBQueue()
	q.Init()
	var wg sync.WaitGroup
	threads := os.Getenv("NUM_THREADS")
	threadsInt, err := strconv.Atoi(threads)
	if err != nil {
		e := fmt.Errorf("error getting thread count:\n%s", err)
		return e
	}
	// FILTER NO 1 - get all tickers from website
	tkrs, err := GetTickers()
	if err != nil {
		e := fmt.Errorf("error getting tickers:\n%s", err)
		return e
	}
	// Distribute load over threads
	// taks per thread
	TPT := len(tkrs) / threadsInt
	// remaining tasks
	RT := len(tkrs) % threadsInt
	startidx := 0
	for i := 0; i < threadsInt; i++ {
		endidx := startidx + TPT
		if i < RT {
			endidx++
		}
		wg.Add(1)
		go divideLoad(&wg, tkrs[startidx:endidx], startidx)
		startidx = endidx
	}
	wg.Wait()
	return nil
}
