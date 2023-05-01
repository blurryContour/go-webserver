package core

import (
	"fmt"
	"sync"
	"time"
)

func Compute(n int, routine bool) string {
	var wg sync.WaitGroup

	// Start computation
	start := time.Now()
	ch := make(chan Result, 4)

	wg.Add(3)
	if routine {
		go Sum(n, ch, &wg)
		go Square(n, ch, &wg)
		go Fibbonacci(n, ch, &wg)
	} else {
		Sum(n, ch, &wg)
		Square(n, ch, &wg)
		Fibbonacci(n, ch, &wg)
	}
	wg.Wait()

	close(ch)
	resultFormatted := ""
	for result := range ch {
		resultFormatted = fmt.Sprintf("%s%s: %d\n", resultFormatted, result.Name, result.Value)
	}
	elapsed := time.Since(start)
	// Finish computation

	content := fmt.Sprintf("%s\n--------------\n%s\nElapsed: %v", "Compute Result", resultFormatted, elapsed)
	return content
}
