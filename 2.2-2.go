package main

import "fmt"

func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y%10, (x+y)%10
	}

	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fib(n))
}
