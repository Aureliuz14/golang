package main

import (
	"bufio"
	"fmt"
	"os"
	"src/model"
	"strconv"
	"strings"
)

var mySlice []model.Capres = make([]model.Capres, 0.0)

func input() {
	var Capres model.Capres
	fmt.Print("1. Nama :")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	Capres.Nama = nama
	fmt.Print("2. Visi :")
	reader2 := bufio.NewReader(os.Stdin)
	visi, _ := reader2.ReadString('\n')
	visi = strings.TrimSpace(visi)
	Capres.Visi = visi
	fmt.Print("3. Misi :")
	reader3 := bufio.NewReader(os.Stdin)
	misi, _ := reader3.ReadString('\n')
	misi = strings.TrimSpace(misi)
	Capres.Misi = misi
	mySlice = append(mySlice, Capres)
}

func list() {
	for i, daftar := range mySlice {
		fmt.Println("Kandidat Ke", i+1)
		fmt.Println("Nama :", daftar.Nama)
		fmt.Println("Visi :", daftar.Visi)
		fmt.Println("Misi :", daftar.Misi)
		fmt.Println("Jumlah Vote :", daftar.Vote)
	}
}

func vote() {
	fmt.Print("Masukkan Nama Capres Pilihan :")
	reader := bufio.NewReader(os.Stdin)
	nilai, _ := reader.ReadString('\n')
	nilai = strings.TrimSpace(nilai)
	for i, value := range mySlice {
		if nilai == value.Nama {
			mySlice[i].Vote++
			println(mySlice[i].Vote)
			break
		}
	}
}

func result() {
	var max = 0
	var equal = 0
	var nama = ""
	for i, daftar := range mySlice {
		fmt.Println("Kandidat Ke", i+1)
		fmt.Println("Nama :", daftar.Nama)
		fmt.Println("Jumlah Vote :", daftar.Vote)
		if daftar.Vote > max {
			max = daftar.Vote
			nama = daftar.Nama
		}
	}
	for _, daftar := range mySlice {
		if daftar.Vote == max {
			equal++
		}
	}
	if equal > 1 {
		fmt.Println("Tidak ada Pemenang")
	} else {
		fmt.Println("Pemenangnya adalah", nama)
	}
}

func judul() {
	fmt.Println("Election Program")
	fmt.Println("================")
}

func menu() {
	fmt.Println("1. Input Candidate")
	fmt.Println("2. List Candidate")
	fmt.Println("3. Vote Candidate")
	fmt.Println("4. Result")
	fmt.Println("5. Exit")
	fmt.Println("[Choose 5 for exit]")
}

func main() {
	for {
		judul()
		menu()
		reader := bufio.NewReader(os.Stdin)
		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)
		value2, _ := strconv.Atoi(value)
		if value2 == 5 {
			fmt.Println("Logout")
			break
		} else if value2 == 1 {
			input()
		} else if value2 == 2 {
			list()
		} else if value2 == 3 {
			vote()
		} else if value2 == 4 {
			result()
		}
	}

}
