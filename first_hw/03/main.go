package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		println("Please, enter a number")
		return
	}
	if num%2 == 0 {
		println("Number is even")
	} else {
		println("Number is odd")
	}
}
