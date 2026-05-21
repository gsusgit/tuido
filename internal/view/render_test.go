package view

import (
	"strings"
	"testing"

	_ "github.com/gsusgit/tuido/internal/i18n/locales"
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
