package scraper

import (
	"log"
	"os"
	"runtime/debug"
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
		log.Println(err)
		debug.PrintStack()
		return err
	}
	tkrs, err := GetTickers()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
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
