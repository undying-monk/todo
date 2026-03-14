package locking

import (
	"sync"
)

func MutexCounter(wg *sync.WaitGroup, mutex *sync.Mutex, x *int32, i int) int32{
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*x++
	// fmt.Println("hello world", i)
	return *x
}

func RaceMutexCounter(x *int32, n int) int32{
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(n)
	for i := range n{
		go MutexCounter(&wg, &mutex,  x, i)
	}
	wg.Wait()
	return *x
}