package main

import (
	"fmt"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go func(wg *sync.WaitGroup, path string) {
			printSiteStatus(path)
			wg.Done()
		}(&wg, path)
	}
	wg.Wait()
}
