package main

import (
	comps "algo-cli/components"
	data "algo-cli/data"
)

func main() {
	comps.Header()
	result := comps.AskText("Give me some info about a question")
	comps.PrintDefaultInfo("Your query is: " + result)
	comps.PrintSearchResults(data.Search(result))
}
