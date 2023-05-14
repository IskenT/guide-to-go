package main

import "fmt"

func main() {
	engineer := Engineer{
		Name: "Alina", 
		Age: 28,
		Project: Project{
			Name: "Build",
			Priority: "High",
			Technologies: []string{"ff", "gg", "hh"},
		},

	}
	fmt.Println(engineer)
	fmt.Printf("%+v\n", engineer)
	fmt.Println(engineer.Name)
	fmt.Println(engineer.Project.Priority)
}

type Engineer struct {
	Name string
	Age int
	Project Project
}

type Project struct {
	Name string
	Priority string 
	Technologies []string
}
