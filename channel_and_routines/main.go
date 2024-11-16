package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.in",
		"http://amazon.in",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _,l := range(links) {
			go checkLink(l, c);
	}

	for i:=0; i<len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link);

	if (err != nil) {
		fmt.Println(link, "might be down!");
		c <- "Might be down I think"
		return;
	}

	fmt.Println(link, "is up");
	c <- "Yep it's up!"

}