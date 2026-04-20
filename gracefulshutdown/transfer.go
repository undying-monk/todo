package main

import (
	"context"
	"fmt"
	"time"
)

func Transfer(ctx context.Context, tasks <-chan int) {
	for {
		// pre-check step to prevent randomness in select
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
		}

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			// db transaction
			fmt.Printf("Started Task %d\n", task)
			time.Sleep(2 * time.Second)
			fmt.Printf("Finished Task %d\n", task)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	tasks := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 1; i <= 3; i++ {
		go Transfer(ctx, tasks)
	}
	time.Sleep(100 * time.Second)
}
