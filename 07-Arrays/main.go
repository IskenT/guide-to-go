package main

import "fmt"

func main() {
	println("Array")
	planets := [8]string{"ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo"}
	var planetsArray [8]string
	planetsArray[0] = "ddd"
	fmt.Println(planets)
	fmt.Println(planetsArray)
	planetSlice := []string{"ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo"}
	fmt.Println(planetSlice)
	var PlanetSliceVerbose []string
	PlanetSliceVerbose = append(PlanetSliceVerbose, "IIOp")
	fmt.Println(PlanetSliceVerbose)

}