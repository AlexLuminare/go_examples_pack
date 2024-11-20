package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // переменная, к которой нужен безопасный доступ
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1) // атомарное увеличение
		}
	}

	// Запускаем несколько горутин
	wg.Add(3)
	go increment()
	go increment()
	go increment()

	wg.Wait()
	fmt.Printf("Итоговое значение счётчика: %d\n", counter) // Ожидается 3000
}
