package main

import (
	"fmt"
	"src/calculator"
	"src/geometry"
)

func main() {
	a := 5.0
	b := 3.0

	fmt.Println(calculator.Divide(a, b))

	r := 14
	fmt.Println(geometry.Lingkaran(float64(r)))

	fmt.Println(geometry.Persegi(a))

	fmt.Println(geometry.Segitiga(a, b))

	fmt.Println(calculator.Multiply(int(a), int(b)))

	fmt.Println(calculator.Add(int(a), int(b)))
}
