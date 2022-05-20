package main

import "fmt"

func Fib(n, m int) int {
	arr := make([]int, 0, 100_000)
	arr = append(arr, 0, 1)

	for i := 0; i < n; i++ {
		arr = append(arr, (arr[i+1]+arr[i])%m)

		if arr[i+2] == 1 && arr[i+1] == 0 {
			period := len(arr) - 2
			return arr[n%period]
		}
	}

	return arr[n] % m
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println(Fib(n, m))
}
