package modules

import (
	"testing"
)

func TestAdd(t *testing.T) {
	add := new(Add)

	if add.Name() != "ADD" {
		t.Error("Invalid Command")
	}

	inputCommand := "8 4"
	var output float64
	output = 12

	if err := add.Validate(""); err == nil {
		t.Error(err)
	}

	if err := add.Validate("a a"); err == nil {
		t.Error(err)
	}

	if err := add.Validate(inputCommand); err != nil {
		t.Error(err)
	}

	result, err := add.Exec()

	if err != nil {
		t.Error(err)
	}

	if result != output {
		t.Error("output is invalid")
	}
}
