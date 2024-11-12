package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args;

	if (len(args) < 2) {
		os.Exit(1);
	}
		
	fileName := args[1];
		
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1);
	}

	os.Stdout.Write(data)

	f, err := os.Open(fileName);

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1);
	}

	io.Copy(os.Stdout, f);
}