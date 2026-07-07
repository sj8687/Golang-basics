package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func increment() {
	defer wg.Done()

	mu.Lock()
	count++
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	fmt.Println(count)
}