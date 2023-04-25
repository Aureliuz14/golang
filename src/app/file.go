package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxnumber(data ...float64) float64 {
	var max float64
	for i, d := range data {
		switch {
		case i == 0:
			max = d
		case d > max:
			max = d
		}
	}
	return max
}

func minnumber(data ...float64) float64 {
	var min float64
	for i, d := range data {
		switch {
		case i == 0:
			min = d
		case d < min:
			min = d
		}
	}
	return min
}

func main11() {
	fmt.Println("Opening file...")
	f, err := os.Open("data.txt")

	if err != nil {
		panic("Sorry, cannot access the file!")
	}

	defer func() {
		fmt.Println("Closing File...")
		f.Close()
	}()

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
		SUM += value
	}
	count := float64(len(data))
	fmt.Println("SUM dari data.txt adalah", SUM)
	AVG := SUM
	SUM /= float64(len(data))
	fmt.Println("AVG dari data.txt adalah", AVG)
	fmt.Println("COUNT dari data.txt adalah", count)
	MAX := maxnumber(data...)
	MIN := minnumber(data...)

	fmt.Println("MIN dari data.txt adalah", MIN)
	fmt.Println("MAX dari data.txt adalah", MAX)

	sum2 := strconv.FormatFloat(SUM, 'f', 3, 64)
	avg2 := strconv.FormatFloat(AVG, 'f', 3, 64)
	count2 := strconv.FormatFloat(count, 'f', 3, 64)
	min2 := strconv.FormatFloat(MIN, 'f', 3, 64)
	max2 := strconv.FormatFloat(MAX, 'f', 3, 64)

	g, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic("Fail To Create File")
	}
	defer func() {
		fmt.Println("Closing File...")
		g.Close()
	}()
	_, err = g.WriteString("SUM :" + sum2 + "\n" + "AVG :" + avg2 + "\n" + "COUNT :" + count2 + "\n" + "MIN :" + min2 + "\n" + "MAX :" + max2 + "\n")
	if err != nil {
		panic("Fail To Write File")
	}
}
