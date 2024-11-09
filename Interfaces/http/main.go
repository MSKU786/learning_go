package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")

	if (err != nil) {
			fmt.Println("Error", err);
			os.Exit(1);
	}

	fmt.Println(resp);

	//Reading data through the reader interface and read function

	bs := make([]byte, 99999)

	resp.Body.Read(bs);

	fmt.Println(string(bs))

}