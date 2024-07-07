package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		println("Please, enter a number >= 0")
		return
	}
	for i := n; i >= 1; i-- {
		println(i)
	}
}
