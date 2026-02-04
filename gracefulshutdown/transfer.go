package main

import (
	"context"
	"fmt"
	"time"
)

func Transfer(ctx context.Context, tasks <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			// db transaction
			time.Sleep(2 * time.Second)
			fmt.Println(task)

			// notification
			// time.Sleep(3 * time.Second)
			// default:
			// 	fmt.Println("default")
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	tasks := make(chan int, 10)
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- i
		}
		close(tasks)
	}()
	go Transfer(ctx, tasks)

	time.Sleep(100 * time.Second)
}
