package main

import (
	"fmt"
	"src/calculator"
	"src/geometry"
	"src/model"
)

func Print(p model.Person) {
	fmt.Println(p.Username, p.Password)
}

func main2() {
	a := 5.0
	b := 3.0

	fmt.Println(calculator.Divide(a, b))

	r := 14
	fmt.Println(geometry.Lingkaran(float64(r)))

	fmt.Println(geometry.Persegi(a))

	fmt.Println(geometry.Segitiga(a, b))

	fmt.Println(calculator.Multiply(int(a), int(b)))

	fmt.Println(calculator.Add(int(a), int(b)))

	// Array

	// username
	// email
	// password
	// string

	var person model.Person

	person.Username = "Michael"
	person.Email = "Michael@gmail.com"
	person.Age = 23
	person.Password = "123456"
	person.Addr.City = "Jambi"
	person.Addr.District = "Jambi Timur"
	person.Addr.SubDistrict = "Tanjung Pinang"
	person.Addr.Street = "Jalan Orang Kayo Hitam No 1000"
	person.Addr.PostalCode = "34146"

	var person2 model.Person
	person2.Username = "Ratu"
	person2.Email = "Ratu@gmail.com"
	person2.Age = 20
	person2.Password = "654321"

	Print(person)
	Print(person2)
}
