package main

import "fmt"

func main() {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		println("Error reading line")
		return
	}
	length := 0
	for range line {
		length += 1
	}
	println(length)
}
