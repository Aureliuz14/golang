package main

import (
	"bufio"
	"fmt"
	"os"
	"src/model"
	"strconv"
	"strings"
)

var date model.Date

func bulanbulan() bool {
	fmt.Print("Masukkan Bulan :")
	reader := bufio.NewReader(os.Stdin)
	bulan, _ := reader.ReadString('\n')
	bulan = strings.TrimSpace(bulan)
	bulan2, _ := strconv.Atoi(bulan)
	b := date.SetMonth(bulan2)
	return b
}

func tahuntahun() {
	fmt.Print("Masukkan Tahun :")
	reader := bufio.NewReader(os.Stdin)
	tahun, _ := reader.ReadString('\n')
	tahun = strings.TrimSpace(tahun)
	tahun2, _ := strconv.Atoi(tahun)
	date.SetYear(tahun2)
	fmt.Println(date.GetYear())
}

func harihari() {
	fmt.Print("Masukkan Hari :")
	reader := bufio.NewReader(os.Stdin)
	hari, _ := reader.ReadString('\n')
	hari = strings.TrimSpace(hari)
	hari2, _ := strconv.Atoi(hari)
	date.SetDay(hari2)
}

func main5() {

	tahuntahun()
	for i := 0; i < 3; i++ {
		b := bulanbulan()
		if b == true {
		}
		for a := 0; a < 3; a++ {
			harihari()
		}
		break
	}

}
