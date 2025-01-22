package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sum := 0
	var mu sync.Mutex

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			mu.Lock()
			sum += square
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Printf("Сумма квадратов: %d\n", sum)
}