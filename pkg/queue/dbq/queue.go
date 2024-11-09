package dbq

import (
	"log"
	"runtime/debug"
	"sync"

	"github.com/k0kubun/pp"
)

type persistence interface {
	Save() error
}

// ======= SINGLY LINKED LIST ==========================================

type Node struct {
	Value persistence
	Next  *Node
}

type sll struct {
	Length int
	Head   *Node
}

func (list *sll) prepend(item persistence) {
	newNode := Node{Value: item, Next: nil}
	if list.Head == nil {
		list.Head = &newNode
		return
	}
	list.Head.Next = &newNode
	list.Head = &newNode
}

func (list *sll) append(item persistence) {
	newNode := Node{Value: item, Next: nil}
	if list.Head == nil {
		list.Head = &newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &newNode
	pp.Println("ADDED:", current.Next)
}

func (list *sll) removeLast() {
	if list.Head == nil {
		return
	}
	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	current := list.Head
	for current.Next != nil && current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (list *sll) removeFirst() {
	if list.Head == nil {
		return
	}
	list.Head = list.Head.Next
	pp.Println("REMOVED:", list.Head)
}

// ======= QUEUE ========================================================

type Queue struct {
	mu          sync.Mutex
	Length      int
	ListOfItems sll
}

func (q *Queue) enqueue(item persistence) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.ListOfItems.append(item)
	q.Length++
}

func (q *Queue) dequeue() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.ListOfItems.Head != nil {
		err := q.ListOfItems.Head.Value.Save()
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}
		q.ListOfItems.removeFirst()
		q.Length--
	}
	return nil
}

func (q *Queue) Enqueue(i persistence) {
	QC <- i
}
