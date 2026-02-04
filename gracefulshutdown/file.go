package main

import (
	"context"
	"fmt"
	"time"
)

func UploadImage(ctx context.Context, image <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case link, ok := <-image:
			if !ok {
				return
			}
			// resize image
			time.Sleep(3 * time.Second)
			fmt.Println(link)

			// upload to cloud storage
			time.Sleep(3 * time.Second)
			// default:
			// 	fmt.Println("default")
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	image := make(chan string)
	// go func() {
	// 	image <- "jpeg"
	// }()
	UploadImage(ctx, image)
}
