package main

import (
	"fmt"
	"net/http"
	"time"
)

func printSiteStatus(path string) {
	var result string
	resp, err := http.Get(path)
	if err != nil {
		result = err.Error()
	} else {
		result = resp.Status
		resp.Body.Close()
	}
	fmt.Printf("%s --> %s\n", path, result)
}

func main() {
	var paths = []string{
		"http://library.nd.edu",
		"http://nd.edu",
	}
	for _, path := range paths {
		go printSiteStatus(path) // HL
	}
	time.Sleep(5 * time.Second)
}
