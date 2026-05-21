package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/theme"
)

func renderThemeView(m *model.Model, t theme.Theme) string {
	width := m.Width
	height := m.Height
	if width < 30 {
		width = 30
	}
	if height < 12 {
		height = 12
	}

	header := renderHeader(m, t, width)
	footer := renderThemeFooter(t, width)

	vpInnerH := BodyHeight(height, header, footer)
	content := buildThemeContent(m, t)
	body := renderPanelBox(width, vpInnerH, content)

	return renderShell(width, height, header, body, footer)
}

func buildThemeContent(m *model.Model, preview theme.Theme) string {
	titleStyle := lipgloss.NewStyle().Foreground(preview.Accent).Bold(true)

	var b strings.Builder
	b.WriteString(titleStyle.Render(panelContentIndent + i18n.T(i18n.KeyThemeTitle)))
	b.WriteString("\n\n")

	for i, th := range theme.Themes {
		b.WriteString(renderThemeRow(th.Name, m.ThemeIdx == i, preview))
		if i < len(theme.Themes)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func renderThemeRow(name string, selected bool, t theme.Theme) string {
	prefix := "  "
	var nameStyle lipgloss.Style
	if selected {
		prefix = "● "
		nameStyle = lipgloss.NewStyle().Foreground(t.Accent).Bold(true)
		if t.SelectedBg != "" {
			nameStyle = nameStyle.Background(t.SelectedBg).Foreground(t.SelectedFg)
		}
	} else {
		nameStyle = lipgloss.NewStyle().Foreground(t.HelpFooter)
	}
	return panelContentIndent + prefix + nameStyle.Render(name)
}

func renderThemeFooter(t theme.Theme, width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	keyList := []struct{ key, desc string }{
		{"↑/↓", i18n.T(i18n.KeyNavigate)},
		{"Enter", i18n.T(i18n.KeyEnterApply)},
		{"Esc", i18n.T(i18n.KeyEscCancel)},
	}
	return renderKeyFooter(t, footerStyle, keyList)
}
