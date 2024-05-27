package main

import (
	"os"

	"github.com/eric3et/smart-cli/functions/common"
	"github.com/eric3et/smart-cli/functions/math"
	"github.com/eric3et/smart-cli/functions/time"
	"github.com/eric3et/smart-cli/types"
	"github.com/eric3et/smart-cli/utils"
)

func main() {

	// debug
	// input := []string{"buffer", "math", "add", "1", "2"}

	// get input arguments
	input := os.Args

	//normalize and categorize input params
	palette, command, args, err := utils.NormalizeInput(input)
	utils.Check(err)

	// get palette definition from input param
	p, err := utils.GetPalette(palette)
	utils.Check(err)

	// get command definition from input param if it exists under chosen palette
	c, err := utils.GetCommand(p, command)
	utils.Check(err)

	// validate arguments provided against command definition
	err = utils.ValidateCommand(c, args)
	utils.Check(err)

	// if command is valid, then execute
	executeCommand(p, c, args)

}

func executeCommand(p *types.Palette, c *types.Command, args []string) {
	switch c.Function {
	case "common.ListCommands":
		common.ListCommands(*p, *c, args)
	case "math.AddNumbers":
		math.AddNumbers(*p, *c, args)
	case "math.MultiplyNumbers":
		math.MultiplyNumbers(*p, *c, args)
	case "time.GetCurrentTime":
		time.GetCurrentTime(*p, *c, args)

	}
}
