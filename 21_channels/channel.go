package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sending

func processNum(numChan chan int){
	for num := range numChan {
	fmt.Println("Processing number", num)
		time.Sleep(time.Second * 2)
}
// fmt.Println("Processing number", <- numChan)

}

func main(){
	// channels are blocking 

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// message := <- messageChan

	// fmt.Println(message) 

	numChan := make(chan int)

	go processNum(numChan)
	// numChan <- 5

	for {
		numChan <- rand.Intn(100)
	}
	// time.Sleep(time.Second * 2)
}