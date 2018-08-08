package modules

import (
	"strconv"
	"strings"

	"github.com/wuriyanto48/clical/api"
)

//Div struct for Div modules
type Div struct {
	Args []string
	x    float64
	y    float64
}

//Name function, returns Name
func (div *Div) Name() string {
	return "DIV"
}

//Validate function
func (div *Div) Validate(args string) error {
	div.Args = strings.Split(args, " ")

	if len(div.Args) < 2 {
		return api.ErrorInvalidArgs
	}

	x, err := strconv.ParseFloat(div.Args[0], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}

	y, err := strconv.ParseFloat(div.Args[1], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}
	div.x = x
	div.y = y
	return nil
}

//Exec returns X / Y
func (div *Div) Exec() (float64, error) {
	if div.y == 0 {
		return 0, api.ErrorDivideByZero
	}

	return div.x / div.y, nil
}
