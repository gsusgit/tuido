package view

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/i18n"
	"github.com/gsus/todo-app/tui/internal/model"
	"github.com/gsus/todo-app/tui/internal/storage"
	"github.com/gsus/todo-app/tui/internal/theme"
)

const prioritySymbol = "●"

type priorityOpt struct {
	p     storage.Priority
	color lipgloss.Color
}

func priorityColor(p storage.Priority, t theme.Theme) lipgloss.Color {
	switch p {
	case storage.PriorityHigh:
		return t.PriorityHigh
	case storage.PriorityLow:
		return t.PriorityLow
	default:
		return t.PriorityMedium
	}
}

func renderPrioritySymbol(p storage.Priority, t theme.Theme) string {
	return lipgloss.NewStyle().Foreground(priorityColor(p, t)).Render(prioritySymbol)
}

func priorityOptions(t theme.Theme) []priorityOpt {
	return []priorityOpt{
		{storage.PriorityHigh, t.PriorityHigh},
		{storage.PriorityMedium, t.PriorityMedium},
		{storage.PriorityLow, t.PriorityLow},
	}
}

func renderPriorityChip(p storage.Priority, label string, selected, focused bool, t theme.Theme) string {
	color := priorityColor(p, t)
	symStr := renderPrioritySymbol(p, t)
	var text string
	if focused && selected {
		text = lipgloss.NewStyle().Foreground(color).Bold(true).Render(label)
	} else if selected {
		text = lipgloss.NewStyle().Foreground(t.Accent).Render(label)
	} else {
		text = lipgloss.NewStyle().Foreground(t.HelpFooter).Render(label)
	}
	return symStr + " " + text
}

func renderPriorityRowFilter(m *model.Model, t theme.Theme) string {
	var parts []string
	allLabel := i18n.T(i18n.KeyAll)
	if m.FilterPriority == "" {
		parts = append(parts, renderPlainChip(allLabel, true, m.FilterCursor == 1, t))
	} else {
		parts = append(parts, renderPlainChip(allLabel, false, m.FilterCursor == 1, t))
	}
	for _, o := range priorityOptions(t) {
		label := m.PriorityLabel(o.p)
		sel := m.FilterPriority == o.p
		parts = append(parts, renderPriorityChip(o.p, label, sel, m.FilterCursor == 1, t))
	}
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += " "
		}
		out += p
	}
	return out
}

func renderPlainChip(label string, selected, focused bool, t theme.Theme) string {
	prefix := ""
	if focused && selected {
		prefix = "● "
	}
	style := lipgloss.NewStyle().Foreground(t.HelpFooter)
	if selected {
		style = style.Foreground(t.Accent)
		if focused {
			style = style.Bold(true)
		}
	}
	return style.Render(prefix + label)
}
