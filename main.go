package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Red("hello")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("hello")

	d := color.New(color.FgBlue, color.Bold)
	d.Println("hello")

	err := "no more hello's"
	red := color.New(color.FgRed).PrintfFunc()
	red("Warning \n")
	red("Error: %s \n", err)

}
