package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct{
	value string
	err error
}

func worker(url string, wg *sync.WaitGroup, resultChan chan result) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)
	fmt.Printf("image processed: %s\n", url)

	resultChan <- result {
		value: url,
		err: nil,
	}
}

func main() {
	var wg sync.WaitGroup

	resultChan := make(chan result, 2)

	startTime := time.Now()

	wg.Add(2)
	go worker("image1.png", &wg, resultChan)
	go worker("image2.png", &wg, resultChan)

	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		fmt.Printf("received: %v \n", result)
	}
	fmt.Printf("It took %s ms.\n", time.Since(startTime))
}
