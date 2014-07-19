package main

import (
	"fmt"

	"github.com/ttacon/goji-colorize/colorize"
)

func main() {
	fmt.Println(colorize.Red, "Writing in colors", colorize.Cyan, "is so much fun", colorize.Reset)
	fmt.Println(colorize.Magenta.Color("You can use colors to color specific phrases"))
	fmt.Printf("%s%s%s\n", colorize.Blue.WithBackground(colorize.White), "And they also have backgrounds!", colorize.Reset)
}
