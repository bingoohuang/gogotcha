package synx

type waiter struct {
	n int
	c chan []Item
}

type state struct {
	items []Item
	wait  []waiter
}

type QueueCh struct {
	s chan state
}

func NewQueueCh() *QueueCh {
	s := make(chan state, 1)
	s <- state{}
	return &QueueCh{s}
}

func (q *QueueCh) Put(item Item) {
	s := <-q.s
	s.items = append(s.items, item)
	for len(s.wait) > 0 {
		w := s.wait[0]
		if len(s.items) < w.n {
			break
		}

		w.c <- s.items[:w.n:w.n]
		s.items = s.items[w.n:]
		s.wait = s.wait[1:]
	}

	q.s <- s
}

func (q *QueueCh) GetMany(n int) []Item {
	s := <-q.s
	if len(s.wait) == 0 && len(s.items) >= n {
		items := s.items[:n:n]
		s.items = s.items[n:]
		q.s <- s
		return items
	}

	c := make(chan []Item)
	s.wait = append(s.wait, waiter{n, c})
	q.s <- s

	return <-c
}
