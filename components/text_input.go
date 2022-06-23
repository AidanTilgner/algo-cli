package components

import (
	"fmt"

	"github.com/pterm/pterm"
)

func GetDefaultTextInput(query string) string {
	printer := &pterm.InteractiveTextInputPrinter{
		DefaultText: query,
		TextStyle:   &pterm.ThemeDefault.PrimaryStyle,
	}
	result, _ := printer.Show()
	fmt.Println()
	return result
}
