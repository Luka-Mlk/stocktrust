package dbq

import (
	"log"
	"runtime/debug"
	"sync"

	"github.com/k0kubun/pp"
)

var QC = make(chan persistence)
var qi = make(chan bool)
var q *Queue

func DBQueue() *Queue {
	if q == nil {
		q = &Queue{mu: sync.Mutex{}}
	}
	return q
}

func (q *Queue) Init() {
	go q.checkerService()
	go q.deQueueService()
}

func (q *Queue) checkerService() {
	for {
		item := <-QC
		q.enqueue(item)
		qi <- true
		pp.Println("ENQUEUE SERVICE")
	}
}

func (q *Queue) deQueueService() {
	for {
		<-qi
		pp.Println("DEQUEUE SERVICE")
		err := q.dequeue()
		if err != nil {
			log.Println(err)
			debug.PrintStack()
		}
	}
}
