package main

import "fmt"

func Gcd(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	if a >= b {
		return Gcd(a%b, b)
	}
	return Gcd(a, b%a)
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(Gcd(a, b))
}
