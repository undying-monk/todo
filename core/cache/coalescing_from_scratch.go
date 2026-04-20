package main

import (
	"fmt"
	"sync"
	"time"
)

type callback struct{
	wg sync.WaitGroup
	val interface{}
	err error
}

var m sync.Mutex
var group map[string]*callback

func doCall(cb *callback, key string, fn func() (interface{}, error) ) {
	cb.val, cb.err = fn()
	m.Lock()
	defer m.Unlock()
	cb.wg.Done()

	if _, ok := group[key]; ok{
		delete(group, key)
	}
}

func Do(key string , fn func() (interface{}, error) ) (v interface{}, err error){
	m.Lock()
	if group == nil{
		group = map[string]*callback{}
	}
	if cb, ok := group[key]; ok{
		m.Unlock()
		cb.wg.Wait()
		return cb.val, cb.err
	}
	cb := new(callback)
	cb.wg.Add(1)
	group[key] = cb
	m.Unlock()

	go doCall(cb, key,fn)
	return cb.val, cb.err
}

var previous interface{}
func simulate(){
	var wg sync.WaitGroup
	// first
	wg.Add(1)
	go func ()  {
		previous = "1"
		fmt.Println("Previous firest", previous)
		wg.Done()
	}()

	// second, do nothing just return val from previous result in memory
	wg.Wait()
	fmt.Println("Previous second", previous)
}

func main(){
	key := "key"
	Do(key, func() (interface{}, error) {
		fmt.Println("first")
		return "", nil
	})
	Do(key, func() (interface{}, error) {
		fmt.Println("second")
		return "", nil
	})

	time.Sleep(2 * time.Second)
}