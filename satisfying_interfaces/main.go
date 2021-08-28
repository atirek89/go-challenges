package main

import (
	"fmt"
	"math/rand"
)

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name   string
	EngAge int
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

// Age function
func (e Engineer) Age() int {
	return e.EngAge
}

func main() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot", EngAge: rand.Intn(50)}
	atirek := Engineer{Name: "Atirek", EngAge: rand.Intn(50)}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
	programmers = append(programmers, atirek)

	fmt.Println(programmers)
}
