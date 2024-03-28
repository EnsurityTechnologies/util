package util

import "sync"

// FIFOQueue represents a first-in-first-out queue.
type FIFOQueue struct {
	queue []interface{}
	lock  sync.Mutex
	cond  *sync.Cond
}

// NewFIFOQueue creates a new FIFO queue.
func NewFIFOQueue() *FIFOQueue {
	q := &FIFOQueue{}
	q.cond = sync.NewCond(&q.lock)
	return q
}

// Push adds an item to the end of the queue.
func (q *FIFOQueue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, item)
	q.cond.Signal() // Signal that a new item is added
}

// Pop removes and returns the first item from the queue.
// If the queue is empty, it waits until an item is available.
func (q *FIFOQueue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.queue) == 0 {
		q.cond.Wait() // Wait until an item is added
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}
