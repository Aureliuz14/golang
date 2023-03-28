package calculator

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}
func Divide(a float64, b float64) float64 {
	if b == 0 {
		return b
	}
	c := a / b
	return c
}
