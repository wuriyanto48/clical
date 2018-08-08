package modules

import (
	"strconv"
	"strings"

	"github.com/wuriyanto48/clical/api"
)

//Mul struct for Mul modules
type Mul struct {
	Args []string
	x    float64
	y    float64
}

//Name function, returns Name
func (mul *Mul) Name() string {
	return "MUL"
}

//Validate function
func (mul *Mul) Validate(args string) error {
	mul.Args = strings.Split(args, " ")

	if len(mul.Args) < 2 {
		return api.ErrorInvalidArgs
	}

	x, err := strconv.ParseFloat(mul.Args[0], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}

	y, err := strconv.ParseFloat(mul.Args[1], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}
	mul.x = x
	mul.y = y
	return nil
}

//Exec returns X * Y
func (mul *Mul) Exec() (float64, error) {
	return mul.x * mul.y, nil
}
