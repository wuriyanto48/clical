package modules

import (
	"strconv"
	"strings"

	"github.com/wuriyanto48/clical/api"
)

//Sub struct for sub modules
type Sub struct {
	Args []string
	x    float64
	y    float64
}

//Name function, returns Name
func (sub *Sub) Name() string {
	return "SUB"
}

//Validate function
func (sub *Sub) Validate(args string) error {
	sub.Args = strings.Split(args, " ")

	if len(sub.Args) < 2 {
		return api.ErrorInvalidArgs
	}

	x, err := strconv.ParseFloat(sub.Args[0], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}

	y, err := strconv.ParseFloat(sub.Args[1], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}
	sub.x = x
	sub.y = y
	return nil
}

//Exec returns X - Y
func (sub *Sub) Exec() (float64, error) {
	return sub.x - sub.y, nil
}
