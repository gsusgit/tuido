package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gsus/todo-app/tui/internal/config"
	"github.com/gsus/todo-app/tui/internal/i18n"
	"github.com/gsus/todo-app/tui/internal/storage"
	"github.com/gsus/todo-app/tui/internal/tuiapp"
)

// RunTUI starts the interactive application.
func RunTUI(cfg config.Config) error {
	i18n.SetLang(cfg.Lang)

	if migrated, err := storage.MigrateFromLegacy(); err != nil {
		return err
	} else if migrated {
		fmt.Println(i18n.T(i18n.KeyMigrating))
	}

	tasks, err := storage.Load()
	if err != nil {
		return fmt.Errorf("load tasks: %w", err)
	}

	prog := tea.NewProgram(tuiapp.NewProgram(cfg, tasks), tea.WithAltScreen())
	if _, err := prog.Run(); err != nil {
		return fmt.Errorf("run tui: %w", err)
	}
	return nil
}
