package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func generator(num ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range num {
			fmt.Println("generator", n)
			ch <- n
		}
	}()
	return ch
}

func receive(n <-chan int, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case x, ok := <-n:
				if !ok {
					return
				}
				ch <- x + 1
			case <-done:
				return
			}
		}
	}()
	return ch
}

func receiveError(n <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 2; i++ {
			v := <-n
			fmt.Println("receiveError", v)
			ch <- v
		}
	}()
	return ch
}

func merge(done <-chan struct{}, out <-chan int, out2 <-chan int) <-chan int {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	do := func(out <-chan int) {
		for v := range out {
			select {
			case ch <- v:
			case <-done:
				return
			}
		}
		wg.Done()
	}

	go do(out)
	go do(out2)
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	ctx := context.Background()
	done := make(chan struct{})
	defer close(done)
	ch := generator(2, 3, 4, 5)
	out := receive(ch, done)
	out2 := receive(ch, done)
	// out := receiveError(ch)
	// fmt.Println("out", <-out)
	// fmt.Println("out", <-out)

	for n := range merge(done, out, out2) {
		fmt.Println(n)
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-sigs
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-done:
	case <-sigs:
		fmt.Println("xx")
	}
}
