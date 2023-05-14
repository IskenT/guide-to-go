package main

import "fmt"

func main() {
	//HelloWorld("World", 27)
	total, negativeN := AddNums(5, 6)
	fmt.Println(total)
	fmt.Println(negativeN) 
}
func HelloWorld(name string, age int) {
	fmt.Println("Hello", name, age)
}

func AddNums (a, b int) (int, int) {
	return a + b, a-b
}

