package main

import (
	comps "algo-cli/components"
	data "algo-cli/data"
	"fmt"
)

func main() {
	comps.Header()
	// result := comps.AskText("What is your name?")
	// comps.PrintDefaultInfo(result)
	fmt.Println(data.Search("Anagram"))
}
