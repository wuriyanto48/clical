package modules

import (
	"strconv"
	"strings"

	"github.com/wuriyanto48/clical/api"
)

//Add struct for add modules
type Add struct {
	Args []string
	x    float64
	y    float64
}

//Name function, returns Name
func (add *Add) Name() string {
	return "ADD"
}

//Validate function
func (add *Add) Validate(args string) error {

	add.Args = strings.Split(args, " ")
	if len(add.Args) < 2 {
		return api.ErrorInvalidArgs
	}

	x, err := strconv.ParseFloat(add.Args[0], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}

	y, err := strconv.ParseFloat(add.Args[1], 64)
	if err != nil {
		return api.ErrorShouldBeNumber
	}
	add.x = x
	add.y = y
	return nil
}

//Exec returns X + Y
func (add *Add) Exec() (float64, error) {
	return add.x + add.y, nil
}
