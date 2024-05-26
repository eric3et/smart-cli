package main

import (
	"fmt"
	"os"

	"github.com/eric3et/smart-cli/functions/math"
	"github.com/eric3et/smart-cli/functions/time"
	"github.com/eric3et/smart-cli/types"
	"github.com/eric3et/smart-cli/utils"
)

func main() {

	//DEBUG
	// input := []string{"buffer", "math", "add", "1", "2"}

	input := os.Args

	palette, command, args := utils.NormalizeInput(input)
	if palette == "" || command == "" {
		fmt.Println("Invalid Request")
	}

	p := utils.GetPalette(palette)
	if p.Name == "" {
		return
	}

	c := utils.GetCommand(p, command)
	if c.Name == "" {
		return
	}
	executeCommand(c, args)

}

func executeCommand(c types.Command, args []string) {
	switch c.Function {
	case "math.AddNumbers":
		math.AddNumbers(c, args)
	case "math.MultiplyNumbers":
		math.MultiplyNumbers(c, args)
	case "time.GetCurrentTime":
		time.GetCurrentTime(c, args)

	}
}
