package synx

import (
	"sync"
)

type Item = int

type Queue struct {
	items []Item
	*sync.Cond
}

func NewQueue() *Queue {
	q := new(Queue)
	q.Cond = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *Queue) Put(item Item) {
	q.L.Lock()
	defer q.L.Unlock()
	q.items = append(q.items, item)
	q.Signal()
}

func (q *Queue) GetMany(n int) []Item {
	q.L.Lock()
	defer q.L.Unlock()
	for len(q.items) < n {
		q.Wait()
	}
	items := q.items[:n:n]
	q.items = q.items[n:]
	return items
}
