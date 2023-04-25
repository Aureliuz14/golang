package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello from website")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerAbout(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Created By Michael")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerSamsung(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Samsung A3 Rp 1.999.000")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerLenovo(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Lenovo 3+  Rp 3.999.000")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerProductSamsung(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	message := []byte("Samsung A5S  Rp 3.499.000" + category[3])
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCalculatorAdd(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	input1 := category[3]
	input2 := category[4]
	conv1, _ := strconv.Atoi(input1)
	conv2, _ := strconv.Atoi(input2)
	hasil := conv1 + conv2
	hasil2 := strconv.Itoa(hasil)
	message := []byte("Hasil dari URL diatas adalah " + hasil2)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCalculatorSubtract(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	input1 := category[3]
	input2 := category[4]
	conv1, _ := strconv.Atoi(input1)
	conv2, _ := strconv.Atoi(input2)
	hasil := conv1 - conv2
	hasil2 := strconv.Itoa(hasil)
	message := []byte("Hasil dari URL diatas adalah " + hasil2)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCalculatorMultiply(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	input1 := category[3]
	input2 := category[4]
	conv1, _ := strconv.Atoi(input1)
	conv2, _ := strconv.Atoi(input2)
	hasil := conv1 * conv2
	hasil2 := strconv.Itoa(hasil)
	message := []byte("Hasil dari URL diatas adalah " + hasil2)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCalculatorDivide(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	input1 := category[3]
	input2 := category[4]
	conv1, _ := strconv.ParseFloat(input1, 64)
	conv2, _ := strconv.ParseFloat(input2, 64)
	hasil := conv1 / conv2
	hasil2 := fmt.Sprintf("%.1f", hasil)
	message := []byte("Hasil dari URL diatas adalah " + hasil2)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCalculator(writer http.ResponseWriter, request *http.Request) {
	category := strings.Split(request.URL.Path, "/")
	operation := category[2]
	input1 := category[3]
	input2 := category[4]
	conv1, _ := strconv.ParseFloat(input1, 64)
	conv2, _ := strconv.ParseFloat(input2, 64)
	hasil := 0.0
	shasil := ""
	if operation == "add" {
		hasil = conv1 + conv2
		shasil = strconv.FormatFloat(hasil, 'f', -1, 64)
	} else if operation == "subtract" {
		hasil = conv1 - conv2
		shasil = strconv.FormatFloat(hasil, 'f', -1, 64)
	} else if operation == "multiply" {
		hasil = conv1 * conv2
		shasil = strconv.FormatFloat(hasil, 'f', -1, 64)
	} else if operation == "divide" {
		hasil = conv1 / conv2
		shasil = strconv.FormatFloat(hasil, 'f', 2, 64)
	}

	message := []byte("<b>Hasil dari URL diatas adalah </b>" + shasil)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerHome(writer http.ResponseWriter, request *http.Request) {
	message := []byte("<b>Welcome To My Website</b>")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCountry(writer http.ResponseWriter, request *http.Request) {
	var arr [5]string
	arr[0] = "Indonesia"
	arr[1] = "Myanmar"
	arr[2] = "Kamboja"
	arr[3] = "Vietnam"
	arr[4] = "Malaysia"

	for i := 0; i < len(arr); i++ {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := fmt.Fprint(writer, "<ul><li>", arr[i], "</li></ul>")

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", handlerAbout)
	http.HandleFunc("/product/samsung", handlerSamsung)
	http.HandleFunc("/product/lenovo", handlerLenovo)
	http.HandleFunc("/product/samsung/", handlerProductSamsung)
	//http.HandleFunc("/calculator/add/", handlerCalculatorAdd)
	//http.HandleFunc("/calculator/subtract/", handlerCalculatorSubtract)
	//http.HandleFunc("/calculator/multiply/", handlerCalculatorMultiply)
	//http.HandleFunc("/calculator/divide/", handlerCalculatorDivide)
	http.HandleFunc("/calculator/", handlerCalculator)
	http.HandleFunc("/home/", handlerHome)
	http.HandleFunc("/country/", handlerCountry)
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
