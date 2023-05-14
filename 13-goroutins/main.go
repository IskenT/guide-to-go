package main

import (
	"fmt"
	"time"
)

func HelloWorld(name string){
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello: %s\n", name )
}
func main() {
	go HelloWorld("World")
	fmt.Println("I will go first")
	time.Sleep(2 * time.Second)
}