package main

import (
	Helpers "./internal/helpers"
	Init "./internal/init"
)

func main() {
	Init.DoBasics()
	Helpers.AssignArgs()
}
