package main

import (
	"fmt"

	"github.com/ttacon/colorize"
)

func main() {
	// You can just use colors
	fmt.Println(colorize.Red, "Writing in colors", colorize.Cyan, "is so much fun", colorize.Reset)
	fmt.Println(colorize.Magenta.Color("You can use colors to color specific phrases"))

	// Or you can use styles
	blueOnWhite := colorize.Blue.NewStyle().WithBackground(colorize.White)
	fmt.Printf("%s%s%s\n", blueOnWhite, "And they also have backgrounds!", colorize.Reset)
	fmt.Println(blueOnWhite.Style("You can style strings the same way you can color them!"))
}
