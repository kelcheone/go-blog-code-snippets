package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		mu.Lock()
		fmt.Println("Goroutine 1 acquired lock")
		mu.Unlock()

		mu.Lock()
		fmt.Println("Goroutine 1 acquired lock again")
		mu.Unlock()

		wg.Done()
	}()

	go func() {
		mu.Lock()
		fmt.Println("Goroutine 2 acquired lock")
		mu.Unlock()

		mu.Lock()
		fmt.Println("Goroutine 2 acquired lock again")
		mu.Unlock()

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")
}
