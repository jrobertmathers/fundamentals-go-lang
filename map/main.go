package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#C70039",
		"green": "#3CFF33",
		"blue":  "#335BFF",
	}

	printMap(colors)

	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
