package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4b657",
		"white": "#ffffff",
	}
	colors["10"] = "#ff780"

	printmap(colors)
}
func printmap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}

}
