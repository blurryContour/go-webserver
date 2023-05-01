package core

import (
	"sync"
	"time"
)

type Result struct {
	Name  string
	Value int
}

func Sum(n int, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := n * (n + 1) / 2
	time.Sleep(time.Millisecond * 100)
	result <- Result{"Sum", sum}
}

func Square(n int, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	square := n * n
	time.Sleep(time.Millisecond * 100)
	result <- Result{"Square", square}
}

func Fibbonacci(n int, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	if n < 2 {
		result <- Result{"Fibbonacci", n}
	}

	time.Sleep(time.Millisecond * 150)

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		x := b
		b = b + a
		a = x
	}

	result <- Result{"Fibbonacci", b}
}
