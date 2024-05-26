package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/eric3et/smart-cli/types"
	"gopkg.in/yaml.v2"
)

func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func NormalizeInput(inputs []string) (palette string, command string, args []string) {

	//allow execution when built or not
	if inputs[0] == "smart" {
		inputs = inputs[1:]
	} else {
		inputs = inputs[0:]
	}

	// convert all inputs to lower case
	for idx, val := range inputs {
		inputs[idx] = strings.ToLower(val)
	}

	// seperate inputs
	palette = inputs[1]
	command = inputs[2]
	args = inputs[3:]

	return palette, command, args
}

func LoadPalettes() (palettes types.Palettes) {
	filename, _ := filepath.Abs("./palettes.yaml")

	yamlFile, err := os.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(yamlFile, &palettes)
	Check(err)

	return palettes
}

func GetPalette(palette string) types.Palette {

	palettes := LoadPalettes()

	// get palette from palettes
	for _, p := range palettes {
		if p.Palette.Name == palette {
			return p.Palette
		}
	}

	fmt.Printf("%s Palette doesn't exist\n", palette)
	return types.Palette{}
}

func GetCommand(palette types.Palette, command string) types.Command {

	// get palette from palettes
	for _, c := range palette.Commands {
		if c.Command.Name == command {
			return c.Command
		}
	}

	fmt.Printf("%s Command doesn't exist in %s Palette\n", command, palette.Name)
	return types.Command{}
}
