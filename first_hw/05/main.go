package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		println("Please, enter a number")
		return
	}
	println(factorial(n))
}
