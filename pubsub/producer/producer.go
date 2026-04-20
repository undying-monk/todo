package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	topic       = flag.String("topic", "my-topic", "topic to consume from")
	group       = flag.String("group", "", "group to consume within")
	logger      = flag.Bool("logger", false, "if true, enable an info level logger")
	seedBrokers = flag.String("brokers", "localhost:9092", "comma delimited list of seed brokers")
)

func main() {
	seeds := []string{"localhost:9092"}
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.DefaultProduceTopic(*topic),
	}
	client, err := kgo.NewClient(opts...)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		client.Produce(ctx, &kgo.Record{
			Topic:     *topic,
			Key:       []byte(strconv.Itoa(i)),
			Value:     []byte(strconv.Itoa(i)),
			Timestamp: time.UnixMilli(int64(i)),
		}, func(r *kgo.Record, err error) {
			defer wg.Done()
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}
		})
	}

	wg.Wait()
}
