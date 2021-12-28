package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Печатает на стандартный вывод и отправляет int в канал
func longRunningProcess(ctx context.Context) error {
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
		err := longRunningProcess(ctx)
		if err != nil {
			cancel()
		}
	}()
	time.Sleep(time.Second)
	afterLongProcess(ctx)
}
