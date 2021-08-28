package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	fmt.Println("Implement Me")

	return Developer{Name: name.(string), Age: age.(int)}
}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Elliot"
	var age interface{} = 26

	s, ok := name.(string)
	fmt.Println(s, ok)

	s1, ok1 := name.(int)
	fmt.Println(s1, ok1)

	// s2 := name.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(s2)

	s2 := name
	fmt.Println(s2)

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
