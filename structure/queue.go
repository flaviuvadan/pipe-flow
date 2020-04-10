// holds the logic of a simple queue data structure that is used by the structure for performing flows
package structure

import "fmt"

// Queue represents a simple queue data structure
type Queue struct {
	size    int              // size of the queue
	objects chan interface{} // objects is a channel of interfaces to which anything can be added
}

// NewQueue creates and returns a new queue data structure
func NewQueue(s int) *Queue {
	return &Queue{
		objects: make(chan interface{}, s),
	}
}

// Push add the given interface to the queue
func (q *Queue) Push(i interface{}) error {
	if len(q.objects) == q.size-1 {
		return fmt.Errorf("queue is full")
	}
	q.objects <- i
	return nil
}

// Pop removes and returns the element at the front of the queue
func (q *Queue) Pop() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return <-q.objects
}

// Size checks and returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.objects)
}
