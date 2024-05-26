package main

import (
	"flag"
	"fmt"
)

func main() {
	out := flag.String("name", "Eric", "")
	flag.Parse()
	fmt.Println("Hi", *out, ", How are you?")
}
