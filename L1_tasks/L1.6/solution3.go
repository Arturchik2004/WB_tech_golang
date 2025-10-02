package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменен, горутина завершается:", ctx.Err())
			return
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(500 * time.Millisecond)
}
