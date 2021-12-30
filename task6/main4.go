package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func longRunningProcess() error {
	time.Sleep(time.Second)
	return errors.New("failed: Too long")
}

func afterLongProcess(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("Last process too long")
	default:
		fmt.Printf("Normal mode")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := longRunningProcess()
		if err != nil {
			//longRunningProcess вернул ошибку, закрываем ctx
			cancel()
		}
	}()
	time.Sleep(time.Second)
	afterLongProcess(ctx)
}
