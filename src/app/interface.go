package main

import (
	"bufio"
	"fmt"
	"os"
	"src/interfaces"
	"src/model"
	"strconv"
	"strings"
)

var screen model.Screen
var file model.File
var pr interfaces.PrintInterfaces

func ShowArea(s interfaces.GeometryInterface) {
	s.ShowArea()
}

func Input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your E-Mail :")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	screen.SetEmail(email)

	fmt.Print("Enter Your Password :")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)
	screen.SetPass(pass)

}

func ShowMenu() {
	fmt.Println("1. Output To Screen ")
	fmt.Println("2. Output To File")
	fmt.Println("3. Exit")
}
func main8() {
	/*
		d := model.Square{}
		d.SetSide(5.0)

		x := model.Triangle{}
		x.SetBase(2.0)
		x.SetHeight(5.0)

		ShowArea(d)
		ShowArea(x)
	*/

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Your E-Mail :")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Enter Your Password :")
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimSpace(pass)

		ShowMenu()
		fmt.Print("Enter Your Choice :")
		reader2 := bufio.NewReader(os.Stdin)
		value, _ := reader2.ReadString('\n')
		value = strings.TrimSpace(value)
		value2, _ := strconv.Atoi(value)
		if value2 == 3 {
			fmt.Println("Logout")
			break
		}
		if value2 == 1 {
			screen.SetEmail(email)
			screen.SetPass(pass)
			pr = screen
		} else if value2 == 2 {
			file.SetEmail(email)
			file.SetPass(pass)
			pr = file
		}
		pr.Output()
	}
}
