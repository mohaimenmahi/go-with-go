package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3) // we can add number of goroutines to wait for in single Add() call, or we can call Add() multiple times just before the goroutines are started
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First goroutine")

		m.Lock()
		score = append(score, 1)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second goroutine")

		m.Lock()
		score = append(score, 2)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third goroutine")

		m.Lock()
		score = append(score, 3)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println(score)
}
