package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var expected int = 5
	var actual int = addition(2, 3)

	if expected != actual {
		t.Errorf("expected value: %v\tvalue obtained: %v", expected, actual)
	}
}

func TestSubtraction(t *testing.T) {
	var expected int = 1
	var actual int = subtraction(3, 2)

	if expected != actual {
		t.Errorf("expected value: %v\tvalue obtained: %v", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	var expected int = 6
	var actual int = multiplication(2, 3)

	if expected != actual {
		t.Errorf("expected value: %v\tvalue obtained: %v", expected, actual)
	}
}

func TestDivision(t *testing.T) {
	var expected int = 2
	var actual int = division(10, 5)

	if expected != actual {
		t.Errorf("expected value: %v\tvalue obtained: %v", expected, actual)
	}
}
