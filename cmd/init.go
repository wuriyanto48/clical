package cmd

import (
	"fmt"
	"strings"

	"github.com/wuriyanto48/clical/api"
	"github.com/wuriyanto48/clical/modules"
)

//Commander wrapper
type Commander struct {
	Calc map[string]api.Command
}

//NewCommander Commander's Constructor returns pointer of Commander
func NewCommander() *Commander {
	calc := make(map[string]api.Command)
	calc["add"] = new(modules.Add)
	calc["sub"] = new(modules.Sub)
	calc["mul"] = new(modules.Mul)
	calc["div"] = new(modules.Div)
	return &Commander{Calc: calc}
}

//Exec function
func (c *Commander) Exec(command string) string {
	cmdSlice := strings.SplitN(command, " ", 2)
	cal, ok := c.Calc[cmdSlice[0]]

	if !ok {
		return "Invalid Command"
	}

	if err := cal.Validate(cmdSlice[1]); err != nil {
		return err.Error()
	}

	res, err := cal.Exec()
	if err != nil {
		return err.Error()
	}

	name := cal.Name()

	result := fmt.Sprintf("%s Result = %f", name, res)

	return result
}
