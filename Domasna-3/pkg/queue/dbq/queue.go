package dbq

import (
	"fmt"
	"sync"
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
	Head *Node
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
			e := fmt.Errorf("error saving from queue:\n%s", err)
			return e
		}
		q.ListOfItems.removeFirst()
		q.Length--
	}
	return nil
}

func (q *Queue) Enqueue(i persistence) {
	qc <- i
}
