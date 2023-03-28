package geometry

// CONST untuk variabel yang tidak bisa diubah
const PHI = 3.14

func Segitiga(a float64, b float64) float64 {
	return 0.5 * a * b
}

func Persegi(a float64) float64 {
	return a * a
}

func Lingkaran(r float64) float64 {
	return PHI * r * r
}
