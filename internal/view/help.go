package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/theme"
)

const controlsPerRow = 3

type controlEntry struct {
	keyLabel string
	descKey  i18n.Key
}

func listControlEntries() []controlEntry {
	return []controlEntry{
		{"↑/k ↓/j", i18n.KeyNavigate},
		{"n", i18n.KeyNewTask},
		{"e", i18n.KeyEdit},
		{"d", i18n.KeyDelete},
		{"space", i18n.KeyComplete},
		{"f", i18n.KeyFilter},
		{"/", i18n.KeySearch},
		{"t", i18n.KeyTheme},
		{"c", i18n.KeyHelpMenu},
		{"esc", i18n.KeyExit},
	}
}

func renderControlCell(e controlEntry, keyStyle, descStyle lipgloss.Style) string {
	return keyStyle.Render(e.keyLabel) + " " + descStyle.Render(i18n.T(e.descKey))
}

func buildControlsContent(t theme.Theme, width int) string {
	titleStyle := lipgloss.NewStyle().Foreground(t.Accent).Bold(true)
	keyStyle := lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	descStyle := lipgloss.NewStyle().Foreground(t.HelpFooter)

	entries := listControlEntries()
	perRow := controlsPerRow
	if width < 50 {
		perRow = 2
	}
	if width < 36 {
		perRow = 1
	}

	var rows []string
	for i := 0; i < len(entries); i += perRow {
		end := i + perRow
		if end > len(entries) {
			end = len(entries)
		}
		cells := make([]string, 0, end-i)
		for _, e := range entries[i:end] {
			cells = append(cells, renderControlCell(e, keyStyle, descStyle))
		}
		rows = append(rows, panelContentIndent+joinParts(cells, footerItemSep))
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render(panelContentIndent + i18n.T(i18n.KeyControls)))
	b.WriteString("\n\n")
	b.WriteString(strings.Join(rows, "\n"))
	return b.String()
}

func renderHelpView(m *model.Model, t theme.Theme) string {
	width := m.Width
	height := m.Height
	if width < 30 {
		width = 30
	}
	if height < 12 {
		height = 12
	}

	header := renderHeader(m, t, width)
	footer := renderHelpFooter(t, width)

	vpInnerH := BodyHeight(height, header, footer)
	content := buildControlsContent(t, width)
	body := renderPanelBox(width, vpInnerH, content)

	return renderShell(width, height, header, body, footer)
}

func renderHelpFooter(t theme.Theme, width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	keyList := []struct{ key, desc string }{
		{"Esc", i18n.T(i18n.KeyBack)},
	}
	return renderKeyFooter(t, footerStyle, keyList)
}
