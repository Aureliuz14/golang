package main

import (
	"bufio"
	"fmt"
	"os"
	"src/model"
	"strconv"
	"strings"
)

type km int

func (input km) toMeter() km {
	return input * 1000
}

func (input km) toMilimeter() km {
	return input * 100000
}

type m int

func (input m) toKilometer() m {
	return input / 1000
}

func (input m) toMilimeter() m {
	return input * 10000
}

func main4() {
	var input1 km
	var input2 m

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("1. Enter Kilometer : ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	ikm, _ := strconv.Atoi(str)

	input1 = km(ikm)
	//fmt.Println(reflect.TypeOf(ikm))
	//fmt.Println(reflect.TypeOf(input1))

	inmeter := input1.toMilimeter()
	fmt.Println("1. In Meter : ", inmeter)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("2. Enter Kilometer : ")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	im, _ := strconv.Atoi(str)

	input2 = m(im)
	inmeter2 := input2.toMilimeter()
	fmt.Println("2. In Meter : ", inmeter2)

	var Student model.Student
	Student.FirstName = "Michael"
	Student.LastName = "Aurelius"
	fmt.Println("Nama Lengkap Dari Nama Diatas :", Student.Namalengkap())

	ss := model.Country{
		Name:    "Indonesia",
		Capital: "Jakarta",
		Province: map[string]string{
			"Jambi":          "Jambi",
			"Sumatera Utara": "Medan",
			"Jawa Barat":     "Bandung",
			"Jawa Timur":     "Surabaya",
		},
	}
	fmt.Println(ss.ProvinceCapital("Jambi"))
	fmt.Println(ss.ProvinceCapital("Jawa Timur"))
	fmt.Println(ss.ProvinceCapital("Sumatera Barat"))
}
