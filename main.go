package main

import (
	comps "algo-cli/components"
	data "algo-cli/data"
)

func main() {
	comps.Header()
	// result := comps.AskText("What is your name?")
	// comps.PrintDefaultInfo(result)
	data.Search("Contains")
}
