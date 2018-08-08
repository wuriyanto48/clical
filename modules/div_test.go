package modules

import (
	"testing"
)

func TestDiv(t *testing.T) {
	div := new(Div)

	if div.Name() != "DIV" {
		t.Error("Invalid Command")
	}

	inputCommand := "8 4"
	var output float64
	output = 2

	if err := div.Validate(""); err == nil {
		t.Error(err)
	}

	if err := div.Validate("a a"); err == nil {
		t.Error(err)
	}

	if err := div.Validate(inputCommand); err != nil {
		t.Error(err)
	}

	result, err := div.Exec()

	if err != nil {
		t.Error(err)
	}

	if result != output {
		t.Error("output is invalid")
	}
}
