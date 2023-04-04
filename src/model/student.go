package model

type Student struct {
	FirstName string
	LastName  string
}

func (input Student) Namalengkap() string {
	return input.FirstName + " " + input.LastName
}
