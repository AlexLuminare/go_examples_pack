package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	bufSize = 10
)

type buffer struct {
	data  []int
	mutex sync.Mutex
	cond  *sync.Cond
}

func newBuffer() *buffer {
	b := &buffer{
		data:  make([]int, 0, bufSize),
		mutex: sync.Mutex{},
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *buffer) put(value int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// Ждем, пока в буфере будет место
	for len(b.data) == bufSize {
		b.cond.Wait()
	}

	b.data = append(b.data, value)
	fmt.Printf("Produced: %d\n", value)

	// Сигнализируем потребителю, что есть новые данные
	b.cond.Signal()
}

func (b *buffer) get() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// Ждем, пока в буфере будут данные
	for len(b.data) == 0 {
		b.cond.Wait()
	}

	value := b.data[0]
	b.data = b.data[1:]
	fmt.Printf("Consumed: %d\n", value)

	// Сигнализируем производителю, что в буфере освободилось место
	b.cond.Signal()

	return value
}

func main() {
	b := newBuffer()

	go func() {
		for i := 0; i < 20; i++ {
			b.put(i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(b.get())
			time.Sleep(time.Millisecond * 200)
		}
	}()

	time.Sleep(time.Second * 3)
}
