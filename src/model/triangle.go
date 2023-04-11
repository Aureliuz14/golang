package model

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (d *Triangle) SetBase(base float64) {
	d.base = base
}

func (d Triangle) GetBase() float64 {
	return d.base
}

func (d *Triangle) SetHeight(height float64) {
	d.height = height
}

func (d Triangle) GetHeight() float64 {
	return d.height
}

func (d Triangle) Area() float64 {
	return d.base * d.height * 1 / 2
}

func (d Triangle) ShowArea() {
	fmt.Println(d.Area())
}
