package dbq

import (
	"log"
	"runtime/debug"
	"sync"
)

var qc = make(chan persistence)
var qi = make(chan bool)
var Q *Queue

func DBQueue() *Queue {
	if Q == nil {
		Q = &Queue{mu: sync.Mutex{}}
	}
	return Q
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
			log.Println(err)
			debug.PrintStack()
		}
	}
}
