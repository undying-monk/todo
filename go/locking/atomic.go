package locking

import (
	"sync"
	"sync/atomic"
)

func AtomicCounter(x *int32, wg *sync.WaitGroup, i int) int32{
	defer wg.Done()
	atomic.AddInt32(x, 1)
	// fmt.Println(*x, i)
	return *x
}

func RaceAtomicCounter(x *int32, n int) int32{
	var wg sync.WaitGroup
	wg.Add(n)
	for i := range n{
		go AtomicCounter(x, &wg, i)
	}
	wg.Wait()
	return *x
}