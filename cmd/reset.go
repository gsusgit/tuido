package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/storage"
)

// RunReset deletes task data with optional confirmation.
func RunReset(force bool) error {
	path, err := storage.DataFilePath()
	if err != nil {
		return err
	}
	if !force {
		fmt.Print(i18n.T(i18n.KeyResetPrompt, path))
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(strings.ToLower(line))
		if line != "y" && line != "yes" && line != "s" && line != "sí" && line != "si" && line != "j" && line != "ja" {
			fmt.Println(i18n.T(i18n.KeyResetCancelled))
			return nil
		}
	}
	if err := storage.Reset(); err != nil {
		return err
	}
	fmt.Println(i18n.T(i18n.KeyResetDone))
	return nil
}
