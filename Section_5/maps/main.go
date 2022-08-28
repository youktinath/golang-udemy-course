package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for col, hex := range c {
		fmt.Println("Hex code for color", col, "is", hex)
	}
}