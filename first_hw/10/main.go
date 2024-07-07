package main

type Rectangle struct {
	width  int
	height int
}

func (self *Rectangle) Field() int {
	return self.width * self.height
}

func main() {
	rect := Rectangle{3, 4}
	println(rect.Field())
}
