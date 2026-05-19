package view

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/theme"
)

// Render draws the current application view.
func Render(m *model.Model) string {
	return RenderViewport(m, renderMainView)
}

func renderMainView(m *model.Model) string {
	t := m.Theme
	switch m.View {
	case model.ViewInputTask:
		return renderInputForm(m, t)
	case model.ViewFilter:
		return renderFilterView(m, t)
	case model.ViewHelp:
		return renderHelpView(m, t)
	default:
		return renderList(m, t)
	}
}

func renderList(m *model.Model, t theme.Theme) string {
	width := m.Width
	height := m.Height

	header := renderHeader(m, t, width)
	footer := renderListFooter(m, t, width, m.Toast)
	panelH := BodyHeight(height, header, footer)

	hintRows := 0
	if listContentRows(m) > panelH {
		hintRows = 2 // gap + hint line
	}
	vpInnerH := panelH - hintRows
	if vpInnerH < 1 {
		vpInnerH = 1
		hintRows = 0
	}

	body := RenderListBody(m, t, width, vpInnerH)
	if hintRows > 0 {
		gap := lipgloss.NewStyle().Width(width).Render("")
		body = lipgloss.JoinVertical(lipgloss.Left, body, gap, renderListScrollHint(m, t, width))
	}

	return renderShell(width, height, header, body, footer)
}
