package main

import "fmt"

func main() {
	//another way to declare a map
	// var colors2 map[string]string
	//or
	// colors3 := make(map[string]string)
	//
	// colors3["white"] = "#ffffff"
	// delete(colors3, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#f4fb00",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, v := range c {
		fmt.Printf("Color %v has hex value %v\n", color, v)
	}
}
