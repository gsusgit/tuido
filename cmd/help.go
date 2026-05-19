package cmd

import (
	"fmt"

	"github.com/gsus/todo-app/tui/internal/i18n"
)

const version = "0.2.0"

// PrintHelp shows CLI usage.
func PrintHelp() {
	fmt.Println(i18n.T(i18n.KeyHelpUsage))
}

// PrintVersion shows version.
func PrintVersion() {
	fmt.Println("tuido", version)
}
