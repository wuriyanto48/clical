package api

import (
	"errors"
)

var (
	//ErrorDivideByZero error type
	ErrorDivideByZero = errors.New("Cannot Divide By Zero")

	//ErrorShouldBeNumber error type
	ErrorShouldBeNumber = errors.New("Values Should Be Number or Digit")

	//ErrorInvalidArgs error type
	ErrorInvalidArgs = errors.New("Invalid Arguments")
)

//Command interface, Generic Calculator command
type Command interface {
	Name() string
	Validate(args string) error
	Exec() (float64, error)
}
