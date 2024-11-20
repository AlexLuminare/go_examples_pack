package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик задач при завершении горутины
	fmt.Printf("Worker %d starting\n", id)

	// Симулируем работу
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Добавляем 3 задачи в группу
	numWorkers := 3
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg) // Запускаем горутины
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All workers finished")
}
