package main

import (
	"fmt"
	"math"
)

func isprime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		println("Please, enter a number")
		return
	}
	for i := 2; i <= n; i++ {
		if isprime(i) {
			println(i)
		}
	}
}
