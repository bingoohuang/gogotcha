package synx

import (
	"fmt"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetMany(n)
			fmt.Printf("%02d: %02d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
}

func TestQueueCh(t *testing.T) {
	q := NewQueueCh()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetMany(n)
			fmt.Printf("%02d: %02d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
}

func BenchmarkQueue(b *testing.B) {
	q := NewQueue()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetMany(n)
			fmt.Printf("%02d: %02d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
}

func BenchmarkQueueCh(b *testing.B) {
	q := NewQueueCh()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetMany(n)
			fmt.Printf("%02d: %02d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
}
