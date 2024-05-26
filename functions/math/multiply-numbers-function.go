package math

import (
	"fmt"
	"strconv"

	"github.com/eric3et/smart-cli/types"
	"github.com/eric3et/smart-cli/utils"
)

func MultiplyNumbers(c types.Command, args []string) {
	out := ""
	result := 1

	for _, v := range args {
		x, err := strconv.Atoi(v)
		result *= x
		utils.Check(err)
	}
	out = fmt.Sprint(result)

	fmt.Println(out)

}
