package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	topic       = flag.String("topic", "my-topic", "topic to consume from")
	group       = flag.String("group", "group2", "group to consume within")
	logger      = flag.Bool("logger", false, "if true, enable an info level logger")
	seedBrokers = flag.String("brokers", "localhost:9092", "comma delimited list of seed brokers")
)

func consume(ctx context.Context, client *kgo.Client) {
	for {
		fetches := client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			fmt.Println("client is closed")
			return
		}
		fetches.EachError(func(t string, p int32, err error) {
			fmt.Fprintf(os.Stderr, "fetch err topic %s partition %d: %v", t, p, err)
			// os.Exit(1) // only for critical error
		})

		var rs []*kgo.Record
		fetches.EachRecord(func(r *kgo.Record) {
			fmt.Println("receive", r.Topic, r.Value)
			rs = append(rs, r)
		})

		if err := client.CommitRecords(ctx, rs...); err != nil {
			fmt.Printf("commit records failed: %v", err)
			continue
		}
	}
}

func main() {
	seeds := []string{"localhost:9092"}
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup(*group),
		kgo.ConsumeTopics(*topic),
		kgo.DisableAutoCommit(),
	}
	client, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	ctx := context.Background()
	fmt.Println("started consumer")
	go consume(ctx, client)

	// do graceful shutdown
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt)

	<-sigs
	fmt.Println("received interrupt signal; closing client")
	done := make(chan struct{})
	go func() {
		defer close(done)
		client.Close()
	}()

	select {
	case <-sigs:
		fmt.Println("received second interrupt signal; quitting without waiting for graceful close")
	case <-done:
	}
}
