package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var instance *Singleton

type Singleton struct {
	data string
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initialized"}
		fmt.Println("Singleton instance created")
	})
	return instance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		s := GetInstance()
		fmt.Println(s.data)
	}()

	go func() {
		defer wg.Done()
		s := GetInstance()
		fmt.Println(s.data)
	}()

	wg.Wait()
}
