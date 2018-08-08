package modules

import (
	"testing"
)

func TestMul(t *testing.T) {
	mul := new(Mul)

	if mul.Name() != "MUL" {
		t.Error("Invalid Command")
	}

	inputCommand := "8 4"
	var output float64
	output = 32

	if err := mul.Validate(""); err == nil {
		t.Error(err)
	}

	if err := mul.Validate("a a"); err == nil {
		t.Error(err)
	}

	if err := mul.Validate(inputCommand); err != nil {
		t.Error(err)
	}

	result, err := mul.Exec()

	if err != nil {
		t.Error(err)
	}

	if result != output {
		t.Error("output is invalid")
	}
}
