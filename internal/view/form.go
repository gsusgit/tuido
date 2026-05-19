package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/i18n"
	"github.com/gsus/todo-app/tui/internal/model"
	"github.com/gsus/todo-app/tui/internal/storage"
	"github.com/gsus/todo-app/tui/internal/theme"
)

func renderInputForm(m *model.Model, t theme.Theme) string {
	width := m.Width
	height := m.Height
	if width < 30 {
		width = 30
	}
	if height < 12 {
		height = 12
	}

	header := renderHeader(m, t, width)
	footer := renderInputFooter(t, width)
	if m.Toast != "" {
		footer = renderFooter(t, width, m.KeyMap, &m.Help, m.Toast)
	}

	vpInnerH := BodyHeight(height, header, footer)
	content := buildInputFormContent(m, t, width)
	body := renderPanelBox(width, vpInnerH, content)

	return renderShell(width, height, header, body, footer)
}

func buildInputFormContent(m *model.Model, t theme.Theme, width int) string {
	contentW := panelContentWidth(width)
	m.TitleInput.Width = contentW

	titleLabelStyle := lipgloss.NewStyle().Foreground(t.Foreground).Bold(true)
	inputStyle := lipgloss.NewStyle().Foreground(t.Foreground).Width(contentW)

	var body strings.Builder
	body.WriteString(titleLabelStyle.Render(panelContentIndent + i18n.T(i18n.KeyTitle)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(inputStyle.Render(m.TitleInput.View()))
	body.WriteString("\n\n")
	body.WriteString(titleLabelStyle.Render(panelContentIndent + i18n.T(i18n.KeyCategory)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderCategoryOptions(m, t))
	body.WriteString("\n\n")
	body.WriteString(titleLabelStyle.Render(panelContentIndent + i18n.T(i18n.KeyPriority)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderPriorityOptions(m, t))

	if m.InputErr != "" {
		errStyle := lipgloss.NewStyle().Foreground(t.PriorityHigh)
		body.WriteString("\n" + panelContentIndent)
		body.WriteString(errStyle.Render(m.InputErr))
	}

	return body.String()
}

func renderCategoryOptions(m *model.Model, t theme.Theme) string {
	var b strings.Builder
	for _, c := range storage.AllCategories() {
		selected := m.InputCategory == c
		prefix := ""
		if m.InputCursor == 1 && selected {
			prefix = "● "
		}
		style := lipgloss.NewStyle().Foreground(t.HelpFooter)
		if selected {
			style = style.Foreground(t.Accent)
			if m.InputCursor == 1 {
				style = style.Bold(true)
			}
		}
		b.WriteString(style.Render(prefix + m.CategoryLabel(c)))
		b.WriteString(" ")
	}
	return b.String()
}

func renderPriorityOptions(m *model.Model, t theme.Theme) string {
	var b strings.Builder
	for _, o := range priorityOptions(t) {
		b.WriteString(renderPriorityChip(o.p, m.PriorityLabel(o.p), m.InputPriority == o.p, m.InputCursor == 2, t))
		b.WriteString(" ")
	}
	return b.String()
}
