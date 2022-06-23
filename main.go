package main

import (
	comps "algo-cli/components"
)

func main() {
	comps.Header()
	result := comps.GetDefaultTextInput("What is your name?")
	comps.PrintDefaultInfo(result)
}
