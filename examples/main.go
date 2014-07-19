package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func main() {
	// You can just use colors
	fmt.Println(chalk.Red, "Writing in colors", chalk.Cyan, "is so much fun", chalk.Reset)
	fmt.Println(chalk.Magenta.Color("You can use colors to color specific phrases"))

	// Or you can use styles
	blueOnWhite := chalk.Blue.NewStyle().WithBackground(chalk.White)
	fmt.Printf("%s%s%s\n", blueOnWhite, "And they also have backgrounds!", chalk.Reset)
	fmt.Println(blueOnWhite.Style("You can style strings the same way you can color them!"))
}
