package main

import (
	"fmt"

	"github.com/mitchellh/colorstring"
)

func main() {
	// using default colors
	colorstring.Println("[blue] Foreground color is blue.")

	colorstring.Println("[_blue_] Background color is blue.")

	colorstring.Println("[_blue_] Background color is blue [reset]and [red] text is red.")

	colorstring.Println("[underline][bold] Text is underlined and bold.")

	colorstring.Println("Password is : [hidden] Poo")

	// instantiating with custom colors
	colorize := colorstring.Colorize{
		Colors: map[string]string{
			"boldblue": "1;34",
			"boldcyan": "1;36",
			"reset":    "0",
		},
	}
	fmt.Println(colorize.Color("[boldblue]Bold Blue Text [reset]and [boldcyan]Bold Cyan Text [reset](Using custom supplied colors)"))

	return
}
