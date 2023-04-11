package model

import "os"

type File struct {
	email string
	pass  string
}

func (d *File) SetEmail(email string) {
	d.email = email
}

func (d File) GetEmail() string {
	return d.email
}

func (d *File) SetPass(pass string) {
	d.pass = pass
}

func (d File) GetPass() string {
	return d.pass
}

func (d File) Output() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic("Fail To Create File")
	}
	_, err = f.WriteString(d.email + " " + d.pass)
	if err != nil {
		panic("Fail To Write File")
	}
	f.Close()
}
