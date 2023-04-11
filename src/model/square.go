package model

import "fmt"

type Square struct {
	side float64
}

func (d *Square) SetSide(side float64) {
	d.side = side
}

func (d Square) Getside() float64 {
	return d.side
}

func (d Square) Area() float64 {
	return d.side * d.side
}

func (d Square) ShowArea() {
	fmt.Println(d.Area())
}
