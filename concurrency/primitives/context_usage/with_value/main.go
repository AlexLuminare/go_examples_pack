package main

import (
	"context"
	"fmt"
)

func main() {
	// Создаём контекст с данными
	ctx := context.WithValue(context.Background(), "userID", 42)

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// Извлекаем данные из контекста
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Printf("Processing request for user ID: %v\n", userID)
	} else {
		fmt.Println("No user ID found in context.")
	}
}
