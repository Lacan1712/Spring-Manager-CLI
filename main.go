package main

import (
	ascii_apresentation "SMC/src/ascii_texts"
	"SMC/src/cmd"
)

func main() {
	ascii_apresentation.Apresentation()
	cmd.Execute()
}
