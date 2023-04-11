package model

import "fmt"

type Screen struct {
	email string
	pass  string
}

func (d *Screen) SetEmail(email string) {
	d.email = email
}

func (d Screen) GetEmail() string {
	return d.email
}

func (d *Screen) SetPass(pass string) {
	d.pass = pass
}

func (d Screen) GetPass() string {
	return d.pass
}

func (d Screen) Print() string { // Boleh dipakai boleh tidak
	return d.email + " " + d.pass // Disarankan jangan dipakai function ini
}

func (d Screen) Output() {
	fmt.Println("E-Mail : ", d.email)
	fmt.Println("Password : ", d.pass)
}
