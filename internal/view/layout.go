package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	panelHeightNumer   = 2 // panel height ≈ 65% of remaining space (2/3)
	panelHeightDenom   = 3
	panelContentIndent = "" // no lateral margin on panel content
	shellTopMargin     = 2 // empty rows between header and panel content
)

func lineCount(s string) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, "\n") + 1
}

// clipToMaxHeight truncates styled content to at most maxH terminal rows.
func clipToMaxHeight(s string, maxH int) string {
	if maxH < 1 {
		maxH = 1
	}
	for lipgloss.Height(s) > maxH {
		lines := strings.Split(s, "\n")
		if len(lines) <= 1 {
			break
		}
		s = strings.Join(lines[:len(lines)-1], "\n")
	}
	return s
}

// padToHeight pads with blank lines until content reaches targetH rows.
func padToHeight(s string, targetH int) string {
	for lipgloss.Height(s) < targetH {
		s += "\n"
	}
	return s
}

// renderShell composes header, body, and footer into exactly height rows.
func renderShell(width, height int, header, body, footer string) string {
	if height < 1 {
		height = 1
	}

	footerH := lipgloss.Height(footer)
	if footerH < 1 {
		footerH = 1
	}

	parts := []string{header}
	for i := 0; i < shellTopMargin; i++ {
		parts = append(parts, "")
	}
	parts = append(parts, body)

	middle := lipgloss.JoinVertical(lipgloss.Left, parts...)
	middleH := lipgloss.Height(middle)
	gap := height - middleH - footerH
	if gap > 0 {
		middle += strings.Repeat("\n", gap)
	}

	out := lipgloss.JoinVertical(lipgloss.Left, middle, footer)
	return lipgloss.NewStyle().Width(width).Render(out)
}

// BodyHeight returns panel inner height (~65% of space between header and footer).
func BodyHeight(height int, header, footer string) int {
	headerH := lipgloss.Height(header)
	footerH := lipgloss.Height(footer)
	available := height - headerH - footerH - shellTopMargin
	if available < 1 {
		available = 1
	}
	h := available * panelHeightNumer / panelHeightDenom
	if h < 1 {
		h = 1
	}
	return h
}
