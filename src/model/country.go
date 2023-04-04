package model

type Country struct {
	Name     string
	Capital  string
	Province map[string]string
}

func (input Country) ProvinceCapital(nama string) string {
	return input.Province[nama]
}
