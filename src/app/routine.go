package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getResponseSize(a string) {
	b, err := http.Get(a)
	if err != nil {
		panic("Can't get to ")
	}
	defer b.Body.Close()
	body, err := ioutil.ReadAll(b.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(body))
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main312() {
	go a()
	go b()
	go getResponseSize("http://www.google.com")
	go getResponseSize("http://www.detik.com")
	go getResponseSize("http://www.shopee.co.id")
	go getResponseSize("http://www.example.com")
	go getResponseSize("http://www.golang.org")
	time.Sleep(1 * time.Second)
}

func main2100() {
	resp, err := http.Get("http://www.detik.com")

	if err != nil {
		panic("Can't get to http://www.detik.com")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(body))

	resp.Body.Close()
}
