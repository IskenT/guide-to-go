package main

import (
	"math/rand"
	"fmt"
	
)

func Calculate(values chan int){
	value := rand.Intn(10)
	fmt.Printf("Number: %d\n", value)
	values <- value
}
func main() {
	values := make(chan int)
	go Calculate(values)
	value := <-values
	fmt.Println(value)
}