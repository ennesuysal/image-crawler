package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	crawler := initCrawler("https://www.yartu.io/")
	err := crawler.crawl()
	if err != nil {
		panic(err)
	}

	wg.Wait()

	counter := 0
	for key, values := range crawler.ImgUrls {
		fmt.Printf("%s:\n", key)
		for i, value := range values {
			fmt.Printf("\t%d.) %s\n", i+1, value)
			counter++
		}
	}
	fmt.Printf("\nCOUNTER: %d\n", counter)
}
