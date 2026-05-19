package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/storage"
	"github.com/gsusgit/tuido/internal/theme"
)

func renderFilterView(m *model.Model, t theme.Theme) string {
	width := m.Width
	height := m.Height
	if width < 30 {
		width = 30
	}
	if height < 12 {
		height = 12
	}

	header := renderHeader(m, t, width)
	footer := renderFilterFooter(t, width, m.HasActiveFilter())

	vpInnerH := BodyHeight(height, header, footer)
	content := buildFilterContent(m, t)
	body := renderPanelBox(width, vpInnerH, content)

	return renderShell(width, height, header, body, footer)
}

func buildFilterContent(m *model.Model, t theme.Theme) string {
	labelStyle := lipgloss.NewStyle().Foreground(t.Foreground).Bold(true)

	var body strings.Builder
	body.WriteString(labelStyle.Render(panelContentIndent + i18n.T(i18n.KeyCategory)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderFilterCategoryRow(m, t))
	body.WriteString("\n\n")
	body.WriteString(labelStyle.Render(panelContentIndent + i18n.T(i18n.KeyPriority)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderPriorityRowFilter(m, t))
	body.WriteString("\n\n")
	body.WriteString(labelStyle.Render(panelContentIndent + i18n.T(i18n.KeyStatus)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderFilterStatusRow(m, t))
	body.WriteString("\n\n")
	body.WriteString(labelStyle.Render(panelContentIndent + i18n.T(i18n.KeySortField)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderFilterSortFieldRow(m, t))
	body.WriteString("\n\n")
	body.WriteString(labelStyle.Render(panelContentIndent + i18n.T(i18n.KeySortDirection)))
	body.WriteString("\n" + panelContentIndent)
	body.WriteString(renderFilterSortDirectionRow(m, t))

	return body.String()
}

func renderFilterCategoryRow(m *model.Model, t theme.Theme) string {
	type opt struct {
		val   storage.Category
		label string
	}
	opts := []opt{
		{"", i18n.T(i18n.KeyAll)},
		{storage.CategoryPersonal, m.CategoryLabel(storage.CategoryPersonal)},
		{storage.CategoryTrabajo, m.CategoryLabel(storage.CategoryTrabajo)},
	}
	var parts []string
	for _, o := range opts {
		sel := m.FilterCategory == o.val
		parts = append(parts, renderPlainChip(o.label, sel, m.FilterCursor == 0, t))
	}
	return strings.Join(parts, "  ")
}

func renderFilterStatusRow(m *model.Model, t theme.Theme) string {
	opts := []struct {
		val   model.FilterStatus
		label string
	}{
		{model.FilterStatusAll, i18n.T(i18n.KeyAll)},
		{model.FilterStatusPending, i18n.T(i18n.KeyPending)},
		{model.FilterStatusCompleted, i18n.T(i18n.KeyStatusCompleted)},
	}
	var parts []string
	for _, o := range opts {
		sel := m.FilterStatus == o.val
		parts = append(parts, renderPlainChip(o.label, sel, m.FilterCursor == 2, t))
	}
	return strings.Join(parts, "  ")
}

func renderFilterSortFieldRow(m *model.Model, t theme.Theme) string {
	fields := []struct {
		val   model.SortField
		label string
	}{
		{model.SortDefault, i18n.T(i18n.KeySortDefault)},
		{model.SortTitle, i18n.T(i18n.KeyTitle)},
		{model.SortPriority, i18n.T(i18n.KeyPriority)},
		{model.SortCategory, i18n.T(i18n.KeyCategory)},
	}
	var parts []string
	for _, f := range fields {
		sel := m.SortField == f.val
		parts = append(parts, renderPlainChip(f.label, sel, m.FilterCursor == 3, t))
	}
	return strings.Join(parts, "  ")
}

func renderFilterSortDirectionRow(m *model.Model, t theme.Theme) string {
	dirs := []struct {
		asc   bool
		label string
	}{
		{false, i18n.T(i18n.KeySortDesc) + " ↓"},
		{true, i18n.T(i18n.KeySortAsc) + " ↑"},
	}
	var parts []string
	for _, d := range dirs {
		sel := m.SortAsc == d.asc
		parts = append(parts, renderPlainChip(d.label, sel, m.FilterCursor == 4, t))
	}
	return strings.Join(parts, "  ")
}
