package main

import (
	"log"
	"math/big"
	"os"
	scraper "stocktrust/pkg/scraper/mse"
	"time"
)

func main() {
	start := time.Now()

	r := new(big.Int)
	log.Println(r.Binomial(1000, 10))

	scraper.Init()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s with %s threads", elapsed, os.Getenv("NUM_THREADS"))
}
