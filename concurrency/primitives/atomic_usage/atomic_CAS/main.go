package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 42
	var wg sync.WaitGroup

	updateValue := func(expected, newValue int32) {
		defer wg.Done()
		for {
			// Попытка обновить значение
			if atomic.CompareAndSwapInt32(&value, expected, newValue) {
				fmt.Printf("Успешно обновлено: %d -> %d\n", expected, newValue)
				break
			} else {
				fmt.Printf("Неудача, текущее значение: %d\n", value)
				break
			}
		}
	}

	wg.Add(2)
	go updateValue(42, 100) // Первая горутина
	go updateValue(42, 200) // Вторая горутина

	wg.Wait()
	fmt.Printf("Итоговое значение: %d\n", value)
}
