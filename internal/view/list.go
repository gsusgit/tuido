package view

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/storage"
	"github.com/gsusgit/tuido/internal/theme"
)

func renderHeader(m *model.Model, t theme.Theme, width int) string {
	counter := renderStatsCounter(m, t)
	counterW := lipgloss.Width(counter)
	if counterW < statsColWidth {
		counterW = statsColWidth
	}

	maxLogoW := width - counterW - 1
	if maxLogoW < 10 {
		maxLogoW = 10
	}

	logo := renderLogo(t, width, m.Height, maxLogoW)
	logoH := lipgloss.Height(logo)
	if logoH < 1 {
		logoH = 1
	}

	spacerW := width - lipgloss.Width(logo)
	if spacerW < counterW {
		spacerW = counterW
	}

	rightCol := lipgloss.NewStyle().
		Width(spacerW).
		Height(logoH).
		Align(lipgloss.Right, lipgloss.Center).
		Render(counter)

	return lipgloss.NewStyle().Width(width).Render(
		lipgloss.JoinHorizontal(lipgloss.Top, logo, rightCol),
	)
}

func renderStatsCounter(m *model.Model, t theme.Theme) string {
	total, completed := m.Stats()
	spin := m.StatsSpinner.View()
	countStyle := lipgloss.NewStyle().Foreground(t.HelpFooter)
	count := countStyle.Render(fmt.Sprintf("%d/%d", completed, total))
	return lipgloss.JoinHorizontal(lipgloss.Center, spin, count)
}

// BuildListContent returns all task lines for the viewport.
func BuildListContent(m *model.Model, t theme.Theme, width int) string {
	var b strings.Builder

	if m.SearchActive {
		searchStyle := lipgloss.NewStyle().Foreground(t.Accent)
		b.WriteString(searchStyle.Render("/ " + m.SearchInput.View()))
		b.WriteString("\n")
	}

	display := m.DisplayTasks()
	if len(display) == 0 {
		emptyStyle := lipgloss.NewStyle().Foreground(t.HelpFooter).Italic(true)
		msg := i18n.T(i18n.KeyEmptyList)
		if m.HasActiveFilter() || (m.SearchActive && m.SearchInput.Value() != "") {
			msg = i18n.T(i18n.KeyEmptyFiltered)
		}
		b.WriteString(emptyStyle.Render(panelContentIndent + msg))
		return b.String()
	}

	for i, task := range display {
		selected := i == m.Cursor
		pendingDelete := task.ID == m.DeletePendingID
		b.WriteString(renderTaskLine(task, selected, pendingDelete, t, width, m))
		if i < len(display)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func renderTaskCheckbox(task storage.Task, selected bool, t theme.Theme) string {
	label := "[ ]"
	if task.Completed {
		label = "[x]"
	}
	style := lipgloss.NewStyle().Foreground(t.HelpFooter)
	if task.Completed {
		style = style.Foreground(t.TaskDoneText)
	}
	if selected {
		style = style.Foreground(t.Accent).Bold(true)
	}
	return style.Render(label)
}

func renderTaskLine(task storage.Task, selected, pendingDelete bool, t theme.Theme, width int, m *model.Model) string {
	checkbox := renderTaskCheckbox(task, selected, t)
	prioStr := renderPrioritySymbol(task.Priority, t)
	catLabel := lipgloss.NewStyle().Foreground(t.HelpFooter).Render("[" + m.CategoryLabel(task.Category) + "]")

	var titleStyle lipgloss.Style
	if task.Completed {
		titleStyle = lipgloss.NewStyle().Foreground(t.TaskDoneText)
	} else if selected {
		titleStyle = lipgloss.NewStyle().Foreground(t.SelectedFg).Bold(true)
	} else {
		titleStyle = lipgloss.NewStyle().Foreground(t.TaskTitle)
	}
	title := titleStyle.Render(task.Title)

	content := panelContentIndent + fmt.Sprintf("%s %s %s %s", checkbox, prioStr, catLabel, title)

	rowStyle := lipgloss.NewStyle().Width(panelInnerWidth(width))
	if task.Completed {
		rowStyle = rowStyle.Faint(true)
	}
	if pendingDelete {
		rowStyle = rowStyle.Foreground(t.PriorityHigh)
	}
	return rowStyle.Render(content)
}

// panelInnerWidth is the content width inside the panel (matches viewport).
func panelInnerWidth(width int) int {
	if width < 1 {
		return 1
	}
	return width
}

// panelContentWidth is the usable text width inside the panel.
func panelContentWidth(width int) int {
	return panelInnerWidth(width)
}

// renderPanelBox sizes panel content to the allotted height (no border).
func renderPanelBox(width, innerH int, content string) string {
	if innerH < 1 {
		innerH = 1
	}
	content = clipToMaxHeight(content, innerH)
	content = padToHeight(content, innerH)

	return lipgloss.NewStyle().
		Width(panelInnerWidth(width)).
		Height(innerH).
		Render(content)
}

// RenderListBody returns the viewport view inside the list panel.
func RenderListBody(m *model.Model, t theme.Theme, width, vpInnerH int) string {
	if vpInnerH < 1 {
		vpInnerH = 1
	}
	m.ListVP.Width = panelInnerWidth(width)
	m.ListVP.Height = vpInnerH
	m.ListBodyH = vpInnerH

	content := BuildListContent(m, t, width)
	m.ListVP.SetContent(content)
	m.SyncViewportOffset()

	return renderPanelBox(width, vpInnerH, m.ListVP.View())
}

func listContentRows(m *model.Model) int {
	tasks := len(m.DisplayTasks())
	if tasks == 0 {
		return 0
	}
	rows := tasks
	if m.SearchActive {
		rows++
	}
	return rows
}

func renderListScrollHint(m *model.Model, t theme.Theme, width int) string {
	arrows := lipgloss.NewStyle().Foreground(t.Accent).Render("↑↓")
	msg := i18n.T(i18n.KeyScrollMore, arrows)
	style := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Italic(true).
		Width(panelInnerWidth(width))

	return style.Render(panelContentIndent + msg)
}
