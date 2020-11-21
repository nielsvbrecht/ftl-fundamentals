package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := calculator.Multiply(2, 0)
	if want == got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivision(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Division(2, 8)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
