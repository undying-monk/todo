package main

import (
	"fmt"
	"sync"
)

func main(){
	// a group of workers
	n := 3
	tasks := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(n)
	results := make(chan int, 10)
	for i := 1; i <= n;i++{
		go func (i int)  {
			defer wg.Done()
			// for v := range tasks{
			// 	fmt.Println(v)
			// }
			for {
				select{
				case x, ok := <- tasks:
					if !ok{
						break
					}
					results <- x
				}
			}
		}(i)
	}

	for i := 1; i <= 10; i++{
		tasks <- i
	}
	close(tasks)
	go func ()  {
		wg.Wait()
		close(results)
	}()
	
	
	for x := range results{
		fmt.Println(x)
	}

	// a pipe of tasks, which many workers will share to process
}