package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

func getLanguages()(string, string, string){
	return "golang", "java", "javascript"
}

func main(){
	result := add(3 ,5)
	fmt.Println(result)

	lanf1, lang2, lang3 := getLanguages()
	fmt.Println(lanf1, lang2, lang3)
}