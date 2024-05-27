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
		fmt.Printf("%s\n", err.Error())
	}
}

func NormalizeInput(inputs []string) (palette string, command string, args []string, err error) {

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

	if len(inputs) < 3 {
		return "", "", nil, fmt.Errorf("not enought arguments")
	}
	// seperate inputs
	palette = inputs[1]
	command = inputs[2]
	args = inputs[3:]

	return palette, command, args, nil
}

func LoadPalettes() (palettes types.Palettes) {
	filename, _ := filepath.Abs("./palettes.yaml")

	yamlFile, err := os.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(yamlFile, &palettes)
	Check(err)

	return palettes
}

func GetPalette(palette string) (*types.Palette, error) {

	palettes := LoadPalettes()

	// get palette from palettes
	for _, p := range palettes {
		if p.Palette.Name == palette {
			return &p.Palette, nil
		}
	}

	return nil, fmt.Errorf("%s Palette doesn't exist", palette)
}

func GetCommand(palette *types.Palette, command string) (*types.Command, error) {

	// get palette from palettes
	for _, c := range palette.Commands {
		if c.Command.Name == command {
			return &c.Command, nil
		}
	}

	return nil, fmt.Errorf("%s Command doesn't exist in %s Palette", command, palette.Name)
}

func ValidateCommand(c *types.Command, args []string) error {
	if len(args) < c.Arguments {
		return fmt.Errorf("not enough arguments")
	}
	return nil
}
