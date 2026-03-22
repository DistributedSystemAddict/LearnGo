package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func main27() {
	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)
	wg3.Add(1)
	go func() {
		err := takingTooLong(ctx)
		if err != nil {
			fmt.Println(err)
			wg3.Done()
			return
		}
		wg3.Done()
	}()
	time.Sleep(2 * time.Second)
	cancel(errors.New("Cancelled by timeout"))
	wg3.Wait()
}

func takingTooLong(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Done!")
		return nil
	case <-ctx.Done():
		fmt.Println("Cancelled!")
		return context.Cause(ctx)
	}
}
