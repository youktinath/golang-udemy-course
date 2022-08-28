package main

import "fmt"

type shape interface{
	getArea() int
}

type square struct{
	length int
}
type triangle struct{
	length int
}

func main() {
	s := square{length: 5}
	t := triangle{length: 6}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() int {
	return 3 * t.length
}

func (s square) getArea() int {
	return 4 * s.length
}