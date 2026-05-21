package cmd

import (
	"fmt"

	"github.com/gsusgit/tuido/internal/i18n"
)

// version is set at link time for releases (-ldflags); default for local builds.
var version = "0.2.0"

// PrintHelp shows CLI usage.
func PrintHelp() {
	fmt.Println(i18n.T(i18n.KeyHelpUsage))
}

// PrintVersion shows version.
func PrintVersion() {
	fmt.Println("tuido", version)
}
