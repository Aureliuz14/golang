package main

import (
	"src/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(2, 3)
	if result != 5 {
		t.Error("add function didnt success")
	}
}

func TestSubtract(t *testing.T) {
	result := calculator.Subtract(2, 3)
	if result != -1 {
		t.Error("add function didnt success")
	}
}

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(2, 3)
	if result != 6 {
		t.Error("add function didnt success")
	}
}

func TestDivide(t *testing.T) {
	result := calculator.Divide(5, 2)
	if result != 2.5 {
		t.Error("add function didnt success")
	}
	result = calculator.Divide(0, 2)
	if result != 0 {
		t.Error("Terserah")
	}
	result = calculator.Divide(2, 0)
	if result != 0 {
		t.Error("Terserah")
	}
}
