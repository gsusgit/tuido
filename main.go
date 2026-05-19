package main

import (
	"fmt"
	"os"

	"github.com/gsusgit/tuido/cmd"
	"github.com/gsusgit/tuido/internal/config"
	"github.com/gsusgit/tuido/internal/i18n"

	_ "github.com/gsusgit/tuido/internal/i18n/locales"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}
	i18n.SetLang(cfg.Lang)

	args := os.Args[1:]
	if len(args) == 0 {
		if err := cmd.RunTUI(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	switch args[0] {
	case "lang":
		code := ""
		if len(args) > 1 {
			code = args[1]
		}
		if err := cmd.RunLang(cfg, code); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "reset":
		force := false
		for _, a := range args[1:] {
			if a == "--force" || a == "-f" {
				force = true
			}
		}
		if err := cmd.RunReset(force); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "-h", "--help", "help":
		cmd.PrintHelp()
	case "-v", "--version", "version":
		cmd.PrintVersion()
	default:
		fmt.Fprintf(os.Stderr, "comando desconocido: %s\n", args[0])
		cmd.PrintHelp()
		os.Exit(1)
	}
}
