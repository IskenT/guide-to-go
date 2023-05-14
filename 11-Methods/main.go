package main

import "fmt"


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

func (e Engineer) Print(){
	fmt.Println("Engineer information: ")
	fmt.Printf("Name: %s\n" , e.Name)
	fmt.Printf("Age: %d\n" , e.Age)
	fmt.Printf("Project: %s\n" , e.Project.Name)
}
func (e *Engineer) UpdateAge(){
	e.Age += 10
}

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
	engineer.Print() 
	engineer.UpdateAge()
	engineer.Print()
}