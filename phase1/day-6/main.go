package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan <- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, j)
		results <- j *2
	}
}

func main() {
	// unbufferd channel
	// fmt.Println("\n Unbuffered channel")

	// ch := make(chan string)
	// go func() {
	// 	ch <- "Hello from goroutine"
	// }()

	// msg := <-ch
	// fmt.Println("Received-----", msg)

	// // bufferd channel
	// fmt.Println("Buffered channel start from here")

	// bufCh := make(chan int, 3)
	// bufCh <- 1
	// bufCh <- 2
	// bufCh <- 3

	// fmt.Println("Buffer full")
	// fmt.Println("Read----", <-bufCh)
	// fmt.Println("Read----", <-bufCh)
	// fmt.Println("Read----", <-bufCh)

	// fmt.Println("\n Select statement start")

	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch1 <- "Message from ch1"
	// }()

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch2 <- "Message from ch1"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case m1 := <-ch1:
	// 		fmt.Println("ch1---", m1)
	// 	case m2 := <-ch2:
	// 		fmt.Println("ch2---", m2)
	// 	}
	// }

	// workers pool
	fmt.Println("Worker pool example")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Println("Result---", <-results)
	}
}
