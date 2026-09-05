package main

import (
	"fmt"
	"sync"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(name, "starting worker")
	fmt.Println(name, "finished working")
}

func main() {

	var wg sync.WaitGroup

	workers := []string{"Golang", "Java", "C++"}

	for _, w := range workers {
		wg.Add(1)
		go worker(w, &wg)
		fmt.Println("Started", w, "from main")
	}

	wg.Wait()

	fmt.Println("Goroutine are finshed")
}