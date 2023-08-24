package main

import "fmt"

func main() {

	//empty map
	//var color map[string]string
	//OR
	color := make(map[string]string)
	//color := map[string]string{
	//	"red":   "laal",
	//	"blue":  "neela",
	//	"green": "hara",
	//}
	color["white"] = "#ffffff"
	color["blue"] = "neela"
	color["green"] = "hara"
	delete(color, "white")

	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color:", color, " Hex:", hex)
	}
}
