package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func main() {
	// You can just use colors
	fmt.Println(chalk.Red, "Writing in colors", chalk.Cyan, "is so much fun", chalk.Reset)
	fmt.Println(chalk.Magenta.Color("You can use colors to color specific phrases"))

	// You can just use text styles
	fmt.Println(chalk.Bold.TextStyle("We can have bold text"))
	fmt.Println(chalk.Underline.TextStyle("We can have underlined text"))
	fmt.Println(chalk.Bold, "But text styles don't work quite like colors :(")

	// Or you can use styles
	blueOnWhite := chalk.Blue.NewStyle().WithBackground(chalk.White)
	fmt.Printf("%s%s%s\n", blueOnWhite, "And they also have backgrounds!", chalk.Reset)
	fmt.Println(
		blueOnWhite.Style("You can style strings the same way you can color them!"))
	fmt.Println(
		blueOnWhite.WithTextStyle(chalk.Bold).
			Style("You can mix text styles with colors, too!"))

	// You can also easily make styling functions thanks to go's functional side
	lime := chalk.Green.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).
		Style
	fmt.Println(lime("look at this cool lime text!"))
}
