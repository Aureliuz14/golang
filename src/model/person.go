package model

type Person struct {
	Username string
	Email    string
	Password string
	Age      int
	Addr     Address
}

type Address struct {
	Street      string
	District    string
	SubDistrict string
	City        string
	PostalCode  string
}
