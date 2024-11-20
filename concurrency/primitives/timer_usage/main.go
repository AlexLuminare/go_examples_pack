package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	// Использование в select
	select {
	case <-timer.C: // блокируется на 2 секунды
		fmt.Println("Таймер сработал")
	case <-time.After(time.Minute):
		fmt.Println("time.After отработал")
	}
	// Остановка таймера
	if !timer.Stop() {
		<-timer.C // очищаем канал если таймер уже сработал
	}

}
