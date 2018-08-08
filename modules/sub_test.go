package modules

import (
	"testing"
)

func TestSub(t *testing.T) {
	sub := new(Sub)

	if sub.Name() != "SUB" {
		t.Error("Invalid Command")
	}

	inputCommand := "8 4"
	var output float64
	output = 4

	if err := sub.Validate(""); err == nil {
		t.Error(err)
	}

	if err := sub.Validate("a a"); err == nil {
		t.Error(err)
	}

	if err := sub.Validate(inputCommand); err != nil {
		t.Error(err)
	}

	result, err := sub.Exec()

	if err != nil {
		t.Error(err)
	}

	if result != output {
		t.Error("output is invalid")
	}
}
