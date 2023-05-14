package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}
type Manager struct{
	Name string
}

func (e *Engineer) GetName() string {
	return "Name: " + e.Name
}
func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}
func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}
func main() {
	engineer := &Engineer{Name: "Elliot"}
	manager := &Manager{Name:"Donn"}
	PrintDetails(engineer)
	PrintDetails(manager)
}