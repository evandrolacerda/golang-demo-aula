package main

import "fmt"

func main() {
	// Arrays
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Slices
	var slice []int = []int{1, 2, 3, 4, 5}

	// Mapas
	var m map[string]int = map[string]int{
		"Maça":    1,
		"Banana":  2,
		"Laranja": 3,
	}

	// Structs
	type Person struct {
		Name string
		Age  int
	}

	var person1 Person = Person{Name: "John", Age: 30}
	var person2 Person = Person{Name: "Jane", Age: 25}

	fmt.Println(numbers)
	fmt.Println(slice)
	fmt.Println(m)
	fmt.Println(person1)
	fmt.Println(person2)
}
