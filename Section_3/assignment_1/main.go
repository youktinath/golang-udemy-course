package main

import "fmt"

func main() {
	zeroToTen := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range zeroToTen {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}