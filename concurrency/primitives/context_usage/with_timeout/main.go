package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаём контекст с таймаутом 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Передаём контекст в функцию
	result := make(chan string)
	go performTask(ctx, result)

	select {
	case res := <-result:
		fmt.Println("Result:", res)
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}

func performTask(ctx context.Context, result chan<- string) {
	// Симулируем длительную задачу
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		// Если контекст завершён, выходим
		fmt.Println("Task canceled:", ctx.Err())
	default:
		// Отправляем результат
		result <- "Task completed successfully!"
	}
}
