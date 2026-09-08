package main

import (
	"fmt"
	"log"
	"strconv"
)

func parseLevel(s string) (int, error) {
	// return patterns (value, err)
	// nil error -> success
	// non nil -> failur

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0 , fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return  0 , fmt.Errorf("level must be 1 and 5")
	} 

	return  n , nil
}

func run () error {
	input := "3"

	level, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("Selected level", level)
	return nil
}


func main() {
	// go dont use exceptions for normal failures
	// functions -> return errors as normal return values

	if err := run(); err != nil {
		log.Fatal(err)
	}

}