package main

import (
	"fmt"
	"sync"
)

func main(){
	workers := make(chan struct{}, 3)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 1; i <= 100;i++{
		go func ()  {
			workers <- struct{}{}
			defer func ()  {
				<- workers
			}()
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}