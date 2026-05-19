package view

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/i18n"
	"github.com/gsus/todo-app/tui/internal/keys"
	"github.com/gsus/todo-app/tui/internal/model"
	"github.com/gsus/todo-app/tui/internal/storage"
	"github.com/gsus/todo-app/tui/internal/theme"
)

const footerItemSep = " "

func styleHelp(h *help.Model, t theme.Theme) {
	h.Styles.ShortKey = lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	h.Styles.ShortDesc = lipgloss.NewStyle().Foreground(t.HelpFooter)
	h.Styles.ShortSeparator = lipgloss.NewStyle().Foreground(t.Border)
	h.ShortSeparator = footerItemSep
}

// renderFooter always shows hotkeys; toast appears right-aligned on the same row.
func renderFooter(t theme.Theme, width int, km keys.ListKeyMap, h *help.Model, toast string) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	styleHelp(h, t)
	if width < 20 {
		width = 20
	}
	h.Width = width
	helpView := h.View(km)
	if helpView == "" {
		helpView = fallbackFooter(t, km)
	}

	row := helpView
	if toast != "" {
		toastStyled := lipgloss.NewStyle().Foreground(t.Accent).Italic(true).Render(toast)
		innerW := width
		if innerW < 10 {
			innerW = 10
		}
		gap := innerW - lipgloss.Width(helpView) - lipgloss.Width(toastStyled)
		if gap < 1 {
			gap = 1
		}
		row = helpView + strings.Repeat(" ", gap) + toastStyled
	}

	return footerStyle.Render(row)
}

// renderListFooter shows hotkeys, or a delete-confirm bar while a task is pending removal.
func renderListFooter(m *model.Model, t theme.Theme, width int, toast string) string {
	if m.DeletePendingID != "" {
		return renderDeleteConfirmFooter(t, width)
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	hotkeyStyle := lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	descStyle := lipgloss.NewStyle().Foreground(t.HelpFooter).Render

	parts := []string{
		hotkeyStyle.Render("n") + " " + descStyle(i18n.T(i18n.KeyNewTask)),
		hotkeyStyle.Render("e") + " " + descStyle(i18n.T(i18n.KeyEdit)),
		hotkeyStyle.Render("d") + " " + descStyle(i18n.T(i18n.KeyDelete)),
		hotkeyStyle.Render("f") + " " + descStyle(i18n.T(i18n.KeyFilter)),
		hotkeyStyle.Render("c") + " " + descStyle(i18n.T(i18n.KeyControls)),
	}
	row := joinParts(parts, footerItemSep)

	if m.HasActiveFilter() {
		filterPart := renderActiveFilterSummary(m, t)
		resetPart := hotkeyStyle.Render("r") + " " + descStyle(i18n.T(i18n.KeyReset))
		row += footerItemSep + filterPart + footerItemSep + resetPart
	}

	row = fitFooterRow(row, width, toast, t)
	return footerStyle.Render(row)
}

// renderActiveFilterSummary builds a compact label for the active filter/sort.
func renderActiveFilterSummary(m *model.Model, t theme.Theme) string {
	var filters []string
	if m.FilterCategory != "" {
		filters = append(filters, m.CategoryLabel(m.FilterCategory))
	}
	if m.FilterPriority != "" {
		filters = append(filters, footerPriorityLabel(m, m.FilterPriority, t))
	}
	if m.FilterStatus != model.FilterStatusAll {
		filters = append(filters, m.CompletionStatusLabel(m.FilterStatus))
	}
	if m.HasCustomSort() {
		filters = append(filters, m.SortFieldLabel()+" "+m.SortDirectionLabel())
	}
	return lipgloss.NewStyle().Foreground(t.Accent).Bold(true).Render(strings.Join(filters, "+"))
}

func footerPriorityLabel(m *model.Model, p storage.Priority, t theme.Theme) string {
	return renderPrioritySymbol(p, t) + m.PriorityLabel(p)
}

func fitFooterRow(row string, width int, toast string, t theme.Theme) string {
	if width < 10 {
		width = 10
	}
	if toast == "" {
		return truncateWidth(row, width)
	}
	toastStyled := lipgloss.NewStyle().Foreground(t.Accent).Italic(true).Render(toast)
	gap := width - lipgloss.Width(row) - lipgloss.Width(toastStyled)
	if gap < 1 {
		maxRow := width - lipgloss.Width(toastStyled) - 1
		if maxRow < 1 {
			maxRow = 1
		}
		row = truncateWidth(row, maxRow)
		gap = 1
	}
	return row + strings.Repeat(" ", gap) + toastStyled
}

func truncateWidth(s string, maxW int) string {
	if maxW < 1 || lipgloss.Width(s) <= maxW {
		return s
	}
	ellipsis := "…"
	for len(s) > 0 && lipgloss.Width(s) > maxW {
		_, size := utf8.DecodeLastRuneInString(s)
		if size == 0 {
			break
		}
		s = s[:len(s)-size]
	}
	for lipgloss.Width(s+ellipsis) > maxW && len(s) > 0 {
		_, size := utf8.DecodeLastRuneInString(s)
		if size == 0 {
			break
		}
		s = s[:len(s)-size]
	}
	return s + ellipsis
}

func fallbackFooter(t theme.Theme, km keys.ListKeyMap) string {
	hotkeyStyle := lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	descStyle := lipgloss.NewStyle().Foreground(t.HelpFooter).Render
	parts := []string{
		hotkeyStyle.Render("n") + " " + descStyle(i18n.T(i18n.KeyNewTask)),
		hotkeyStyle.Render("e") + " " + descStyle(i18n.T(i18n.KeyEdit)),
		hotkeyStyle.Render("d") + " " + descStyle(i18n.T(i18n.KeyDelete)),
		hotkeyStyle.Render("c") + " " + descStyle(i18n.T(i18n.KeyControls)),
	}
	return joinParts(parts, footerItemSep)
}

func renderDeleteConfirmFooter(t theme.Theme, width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	hotkeyStyle := lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	descStyle := lipgloss.NewStyle().Foreground(t.HelpFooter).Render

	question := lipgloss.NewStyle().Foreground(t.PriorityHigh).Bold(true).Render(i18n.T(i18n.KeyDelete) + "?")
	keys := joinParts([]string{
		hotkeyStyle.Render("Enter") + " " + descStyle(i18n.T(i18n.KeyDelete)),
		hotkeyStyle.Render("Esc") + " " + descStyle(i18n.T(i18n.KeyEscCancel)),
	}, footerItemSep)

	row := question + footerItemSep + keys
	return footerStyle.Render(truncateWidth(row, width))
}

func renderInputFooter(t theme.Theme, width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	keyList := []struct{ key, desc string }{
		{"Tab", i18n.T(i18n.KeyTabNext)},
		{"Enter", i18n.T(i18n.KeyEnterSave)},
		{"Esc", i18n.T(i18n.KeyEscCancel)},
	}
	return renderKeyFooter(t, footerStyle, keyList)
}

func renderFilterFooter(t theme.Theme, width int, filterActive bool) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Width(width)

	keyList := []struct{ key, desc string }{
		{"Tab", i18n.T(i18n.KeyTabNext)},
		{"Enter", i18n.T(i18n.KeyEnterApply)},
	}
	if filterActive {
		keyList = append(keyList, struct{ key, desc string }{"r", i18n.T(i18n.KeyReset)})
	}
	keyList = append(keyList, struct{ key, desc string }{"Esc", i18n.T(i18n.KeyEscCancel)})
	return renderKeyFooter(t, footerStyle, keyList)
}

func renderKeyFooter(t theme.Theme, footerStyle lipgloss.Style, keyList []struct{ key, desc string }) string {
	hotkeyStyle := lipgloss.NewStyle().
		Background(t.HotkeyBg).
		Foreground(t.HotkeyFg).
		Bold(true).
		Padding(0, 1)
	descStyle := lipgloss.NewStyle().Foreground(t.HelpFooter).Render

	var parts []string
	for _, k := range keyList {
		parts = append(parts, hotkeyStyle.Render(k.key)+" "+descStyle(k.desc))
	}
	return footerStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, joinParts(parts, footerItemSep)))
}

func joinParts(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	out := parts[0]
	for i := 1; i < len(parts); i++ {
		out += sep + parts[i]
	}
	return out
}
