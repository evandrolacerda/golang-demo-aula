package main

import "fmt"

func main() {
	// Números inteiros
	var numInt int = 10
	var numInt8 int8 = 127
	var numInt16 int16 = 32767
	var numInt32 int32 = 2147483647
	var numInt64 int64 = 9223372036854775807

	// Números em ponto flutuante
	var numFloat32 float32 = 3.14
	var numFloat64 float64 = 3.141592653589793

	// Números complexos
	var numComplex64 complex64 = 3 + 4i
	var numComplex128 complex128 = 3.14 + 2.71i

	// Booleanos
	var isTrue bool = true

	// Caracteres Unicode
	var char rune = 'A'

	// Strings
	var text string = "Hello, Golang!"

	fmt.Println(numInt, numInt8, numInt16, numInt32, numInt64)
	fmt.Println(numFloat32, numFloat64)
	fmt.Println(numComplex64, numComplex128)
	fmt.Println(isTrue)
	fmt.Println(char)
	fmt.Println(text)
}
