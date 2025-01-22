package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []string{"a", "b", "c", "d", "e"}
	result := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, value := range data {
		wg.Add(1)
		go func(index int, val string) {
			defer wg.Done()
			mu.Lock()
			result[val] = index
			mu.Unlock()
		}(i, value)
	}

	wg.Wait()
	fmt.Println("Результат:", result)
}