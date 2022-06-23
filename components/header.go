package components

import (
	s "algo-cli/styles"

	p "github.com/pterm/pterm"
)

func Header() {
	header := p.DefaultHeader.WithBackgroundStyle(s.BlueBackground()).WithFullWidth(true)
	p.DefaultCenter.Println(header.Sprint("Welcome to Algo CLI"))
}
