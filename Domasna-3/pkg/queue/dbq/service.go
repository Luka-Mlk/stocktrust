package dbq

import (
	"log"
	"sync"
)

var qc = make(chan persistence)
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
		item := <-qc
		q.enqueue(item)
		qi <- true
	}
}

func (q *Queue) deQueueService() {
	for {
		<-qi
		err := q.dequeue()
		if err != nil {
			log.Printf("error dequing:\n%s", err)
		}
	}
}
