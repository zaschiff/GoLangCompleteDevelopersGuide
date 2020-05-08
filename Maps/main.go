package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("hex code for", key, "is", value)
	}
}

/*
*************************************************** NOTES ****************************************************
	Maps are like dictionaries in python. They a collection of key-value pairs.
	Keys must be the same type to one another as do values.

	syntax for creating a map: Using the code color := map[string]string{"red":"#ff0000"}
		colors    -> variable name
		map       -> map type
		[string]  -> key data type
		string    -> value data type
		{}        -> the map body type collection
		"red"     -> key
		"#ff0000" -> value

	three ways to making a map:
		1. colors := map[string]string{key:value}
		2. var colros map[string]string
		3. colors := make(map[string]string)
*/
