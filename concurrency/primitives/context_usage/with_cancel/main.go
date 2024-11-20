package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go monitor(ctx)

	time.Sleep(2 * time.Second) // Симулируем некоторую работу
	cancel()                    // Отменяем контекст
	time.Sleep(1 * time.Second)
}

func monitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Monitoring stopped:", ctx.Err())
			return
		default:
			fmt.Println("Monitoring in progress...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
