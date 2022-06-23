package components

import (
	t "algo-cli/types"
	"fmt"

	p "github.com/pterm/pterm"
)

func PrintDefaultInfo(info string) {
	p.Info.Printfln(info)
}

func PrintSearchResults(questions []t.Question) {
	fmt.Println("\nSearch Results: ")
	for i, q := range(questions) {
		p.Info.Printfln("%v: %s", i + 1, q.Name)
	}
}