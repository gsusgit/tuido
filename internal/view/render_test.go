package view

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	_ "github.com/gsusgit/tuido/internal/i18n/locales"
	"github.com/muesli/termenv"
	"github.com/gsusgit/tuido/internal/config"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/storage"
	"github.com/gsusgit/tuido/internal/theme"
)

func testModel(w, h int) *model.Model {
	m := model.New(config.Config{})
	m.Width = w
	m.Height = h
	m.Theme = theme.Themes[0]
	return &m
}

func TestBuildListContentEmpty(t *testing.T) {
	m := testModel(80, 24)
	out := BuildListContent(m, m.Theme, 80)
	if !strings.Contains(out, "Press") && !strings.Contains(out, "n") {
		// empty list message (en) or key fallback
		if strings.TrimSpace(out) == "" {
			t.Fatal("expected non-empty empty-list message")
		}
	}

	m.Tasks = []storage.Task{{ID: "1", Title: "hidden"}}
	m.FilterCategory = storage.CategoryPersonal
	m.Tasks[0].Category = storage.CategoryTrabajo
	got := BuildListContent(m, m.Theme, 80)
	if strings.Contains(got, "hidden") {
		t.Fatal("filtered empty should not show non-matching task")
	}
}

func TestRenderSmoke(t *testing.T) {
	m := testModel(80, 24)
	m.Tasks = []storage.Task{{ID: "1", Title: "Task one"}}

	if out := Render(m); strings.TrimSpace(out) == "" {
		t.Fatal("Render returned empty output")
	}

	m.View = model.ViewTheme
	if out := Render(m); strings.TrimSpace(out) == "" {
		t.Fatal("theme view render returned empty")
	}
}

func TestSelectedTaskTitleUsesForeground(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)

	// Omarchy selection_foreground is dark (for light selection chips), not for list rows.
	tm := theme.Theme{
		Foreground: lipgloss.Color("#bcc1dc"),
		SelectedFg: lipgloss.Color("#111422"),
		TaskTitle:  lipgloss.Color("#8899aa"),
		Accent:     lipgloss.Color("#69c3ff"),
		HelpFooter: lipgloss.Color("#555555"),
	}
	task := storage.Task{ID: "1", Title: "Pick me"}
	m := testModel(80, 24)

	sel := renderTaskLine(task, true, false, tm, 80, m)
	if !strings.Contains(sel, "188;193;220") {
		t.Fatalf("selected title should use Foreground (#bcc1dc), got %q", sel)
	}
	if strings.Contains(sel, "17;20;34") {
		t.Fatal("selected title should not use dark SelectedFg (#111422)")
	}
}
