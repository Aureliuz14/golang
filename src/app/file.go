package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Opening file...")
	f, err := os.Open("data.txt")

	if err != nil {
		panic("Sorry, cannot access the file!")
	}

	var data []float64
	scanner := bufio.NewScanner(f)
	SUM := 0.0
	for scanner.Scan() == true {
		val := scanner.Text()
		val2, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic("file content is not a number!")
		}

		data = append(data, val2)

	}
	for _, value := range data {
		a := SUM + value
		fmt.Println(a)
	}
}
