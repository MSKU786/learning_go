package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	for _,l := range(links) {
			checkLink(l);
	}
}

func checkLink(link string) {

	_, err := http.Get(link);

	if (err != nil) {
		fmt.Println(link, "might be down!");
		return;
	}

	fmt.Println(link, "is up");
}