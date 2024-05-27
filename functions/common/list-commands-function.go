package common

import (
	"fmt"

	"github.com/eric3et/smart-cli/types"
)

func ListCommands(p types.Palette, c types.Command, args []string) {
	out := ""
	if len(args) != c.Arguments {
		out = "Incorrect number of arguments"
	}
	out += fmt.Sprintf("%s command list:\n", p.Name)
	for _, v := range p.Commands {
		out += fmt.Sprintf("\t* %s\t---- %s\n", v.Command.Name, v.Command.Help)
	}

	fmt.Println(out)
}
