package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var isShuttingDown atomic.Bool

const (
	_readinessDelay = 5 * time.Second
	_shutdownTimeout = 10 * time.Second
	_shutdownHardPeriod = 3 * time.Second
)

func main() {
	server := http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if isShuttingDown.Load(){
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Shutting down"))
			return
		}
		w.WriteHeader(http.StatusOK)
    	w.Write([]byte("ok"))
	})
	ctx := context.Background()

	fmt.Println("Server running at http://localhost:3000/events")
	go func ()  {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// when parent ctx is done or signal arrived, so return context is done
	rootCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop() // release resource and reset behavior

	<- rootCtx.Done() // wait for context is deadline
	isShuttingDown.Store(true)

	// give time for readiness check, should greater than interval readiness
	time.Sleep(_readinessDelay)

	// drain inflight requests, with shutdown timeout
	shutdowCtx, cancel := context.WithTimeout(ctx, _shutdownTimeout)
	defer cancel()

	// db close, cache close, message broker close, ...

	if err := server.Shutdown(shutdowCtx); err != nil{
		if errors.Is(err, context.DeadlineExceeded){
			time.Sleep(_shutdownHardPeriod)
			fmt.Println("shutdown is timeout")
		}
	}

	fmt.Println("Graceful shutdown successfully")
}