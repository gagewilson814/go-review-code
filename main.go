package main

import "fmt"

func a(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a = a + b
		b = b + 1
	}
}

func fibonacci(ch chan int, n int) {
	a, b := 0, 1
	for i := 0; i <= n; i++ {
		ch <- a
		a, b = a+b, a
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	n := 10
	go fibonacci(ch2, n)
	go a(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Fibonacci sequence:")
	for num := range ch2 { // iterate through the channel
		fmt.Println(num)
	}
}
