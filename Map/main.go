package main

import "fmt"

func main () {

	/*  How to intialize an empty map both the methods will intialize an empty map

		colors := make(map[string]string)
		var colors map[string]string

	*/
		colors := map[string] string {
			"red" : "ff0000",
			"green": "#4fbf45",
		}

		// add a value in map
		colors["blue"] = "#00ff00"

		// delete a value in map
		//delete(colors, "green");

		printMap(colors);
}


func printMap(c map[string]string) {
		for color, hex := range(c) {
			fmt.Println("The hex code for color", color, "is", hex)
		}
}