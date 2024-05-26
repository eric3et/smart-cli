package math

import (
	"fmt"
	"strconv"

	"github.com/eric3et/smart-cli/types"
	"github.com/eric3et/smart-cli/utils"
)

func AddNumbers(c types.Command, args []string) {
	out := ""
	result := 0
	if len(args) != c.Arguments {
		out = "Incorrect number of arguments"
	}

	for _, v := range args {
		x, err := strconv.Atoi(v)
		result += x
		utils.Check(err)
	}
	out = fmt.Sprint(result)

	fmt.Println(out)
}
