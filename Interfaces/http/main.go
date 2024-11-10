package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter  struct{}

func main() {
	resp, err := http.Get("https://www.google.com")

	if (err != nil) {
			fmt.Println("Error", err);
			os.Exit(1);
	}

	fmt.Println(resp);

	//Reading data through the reader interface and read function

	// First way to output data using reader interface
	bs := make([]byte, 99999)

	resp.Body.Read(bs);

	fmt.Println(string(bs))

	lw := logWritter{};

	//Using wirte  2nd way to use io copy function to use write interface interally
	io.Copy(os.Stdout, resp.Body);

	io.Copy(lw, resp.Body);
}



//Customer writer

func (logWritter) Write (bs []byte) (int , error) {
		fmt.Println(string(bs));
		fmt.Println("just wrote this many bytes", len(bs))
		return len(bs), nil;
}