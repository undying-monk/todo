package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var requestGroup singleflight.Group

func getData(ctx context.Context, key string, timeout time.Duration) (interface{}, error){
	res1c := requestGroup.DoChan(key, func ()(interface{}, error)  {
		time.Sleep(timeout)
		return "", nil
	})

	select {
		case <- ctx.Done():
			requestGroup.Forget(key)
			return "", ctx.Err()
		case res := <- res1c:
			if res.Err != nil{
				return "", res.Err
			}
			return res.Val, res.Err
	}
}

func main(){
	ctx := context.Background()


	timeouts := []time.Duration{
		500 * time.Millisecond, // Will time out
		800 * time.Millisecond, // Will time out
		1500 * time.Millisecond, // Will time out
		2500 * time.Millisecond, // WILL SUCCEED
		3000 * time.Millisecond, // WILL SUCCEED
	}
	
	var wg sync.WaitGroup
	for i := 0; i < len(timeouts); i++{
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()

			timeoutCtx, cancel := context.WithTimeout(ctx, timeouts[i])
			defer cancel()

			res, err := getData(timeoutCtx, "key", 2 * time.Second)
			fmt.Println(i, res, err)
		}(i)
	}
	wg.Wait()
}