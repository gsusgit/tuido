package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/config"
	"github.com/gsus/todo-app/tui/internal/i18n"
)

type langModel struct {
	choices []i18n.Lang
	cursor  int
	cfg     config.Config
	quitting bool
}

func (m langModel) Init() tea.Cmd { return nil }

func (m langModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			sel := m.choices[m.cursor]
			m.cfg.Lang = sel.Code
			_ = config.Save(m.cfg)
			i18n.SetLang(sel.Code)
			fmt.Println(i18n.T(i18n.KeyLangChanged, sel.Name))
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m langModel) View() string {
	if m.quitting {
		return ""
	}
	var b strings.Builder
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#89b4fa"))
	b.WriteString(title.Render(i18n.T(i18n.KeySelectLang)))
	b.WriteString("\n\n")
	for i, c := range m.choices {
		prefix := "  "
		if i == m.cursor {
			prefix = "● "
		}
		line := c.Name
		if i == m.cursor {
			line = lipgloss.NewStyle().Bold(true).Render(line)
		}
		b.WriteString(prefix + line + "\n")
	}
	b.WriteString("\n" + i18n.T(i18n.KeyLangNav))
	return b.String()
}

// RunLang sets or interactively selects language.
func RunLang(cfg config.Config, code string) error {
	i18n.SetLang(cfg.Lang)
	if code != "" {
		if !i18n.IsSupported(code) {
			fmt.Fprintln(os.Stderr, i18n.T(i18n.KeyResetInvalid))
			return fmt.Errorf("unsupported language: %s", code)
		}
		cfg.Lang = code
		if err := config.Save(cfg); err != nil {
			return err
		}
		for _, l := range i18n.Available() {
			if l.Code == code {
				fmt.Println(i18n.T(i18n.KeyLangChanged, l.Name))
				break
			}
		}
		return nil
	}
	p := tea.NewProgram(langModel{
		choices: i18n.Available(),
		cfg:     cfg,
	})
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
