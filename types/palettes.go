package types

type Palettes []struct {
	Palette Palette `yaml:"palette"`
}
type Command struct {
	Name      string      `yaml:"name"`
	Function  string      `yaml:"function"`
	Help      string      `yaml:"help"`
	Arguments int         `yaml:"arguments"`
	Default   interface{} `yaml:"default"`
}
type Commands struct {
	Command Command `yaml:"command"`
}
type Palette struct {
	Name     string     `yaml:"name"`
	Commands []Commands `yaml:"commands"`
}
