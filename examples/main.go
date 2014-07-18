package main

import (
	"fmt"
	"github.com/ttacon/goji-colorize/colorize"
)

func main() {
	fmt.Println(colorize.Red, "hello", colorize.Cyan, "world", colorize.Reset)
	fmt.Println(colorize.Magenta.Color("hello, world"))

	fmt.Printf("%s%s%s\n", colorize.Blue.WithBackground(colorize.White), "Hello, World!", colorize.Reset)
}
