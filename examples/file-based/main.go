package main

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

func main() {
	f, err := os.Create("coolio")
	if err != nil {
		fmt.Println("failed to open file:", err)
		return
	}

	chalk.DetectTerminal(f.Fd())

	text := chalk.Red.Color("YOLO")
	f.Write([]byte(text))
	fmt.Println("text: ", text)
	f.Close()
}
