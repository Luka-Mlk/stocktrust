package scraper

import (
	"log"
	"runtime/debug"
	"stocktrust/pkg/queue/dbq"
	"sync"
)

func Init() error {
	q := dbq.DBQueue()
	q.Init()
	var wg sync.WaitGroup
	threads := 10
	tkrs, err := GetTickers()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	TPT := len(tkrs) / threads
	RT := len(tkrs) % threads
	startidx := 0
	for i := 0; i < threads; i++ {
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
