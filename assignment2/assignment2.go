package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args;
	fileName := args[2];

	// if (fileName == nil) {
	// 	os.Exit(1)
	// }
		
	fmt.Println(fileName)
}