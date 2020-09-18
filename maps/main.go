package main

import "fmt"

type col map[string]string

func main() {

	// option 1
	// var colors map[string]string

	// option 2
	//colors := make(map[string]string)

	// options 3
	colors := col{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	colors.add("new")
	printMap(colors)
}

func (colors col) add(c string) {
	colors[c] = c
}

func printMap(c map[string]string) {
	for _, v := range c {
		fmt.Println(v)
	}
}
