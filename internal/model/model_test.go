package model

import (
	"testing"

	"github.com/gsusgit/tuido/internal/config"
	"github.com/gsusgit/tuido/internal/storage"
	"github.com/gsusgit/tuido/internal/theme"
)

func TestDisplayTasksFilter(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{
		{ID: "1", Title: "A", Category: storage.CategoryPersonal, Priority: storage.PriorityHigh},
		{ID: "2", Title: "B", Category: storage.CategoryTrabajo, Priority: storage.PriorityLow},
	}
	m.FilterCategory = storage.CategoryPersonal
	got := m.DisplayTasks()
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("expected 1 personal task, got %d", len(got))
	}
}

func TestDisplayTasksFilterByStatus(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{
		{ID: "1", Title: "Open", Completed: false},
		{ID: "2", Title: "Done", Completed: true},
	}
	m.FilterStatus = FilterStatusPending
	got := m.DisplayTasks()
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("expected 1 pending task, got %d", len(got))
	}
	m.FilterStatus = FilterStatusCompleted
	got = m.DisplayTasks()
	if len(got) != 1 || got[0].ID != "2" {
		t.Fatalf("expected 1 completed task, got %d", len(got))
	}
}

func TestDeletePending(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{{ID: "x", Title: "t"}}
	pending, deleted := m.DeleteByDisplayIndex(0)
	if !pending || deleted {
		t.Fatal("first delete should be pending only")
	}
	_, deleted = m.DeleteByDisplayIndex(0)
	if !deleted || len(m.Tasks) != 0 {
		t.Fatal("second delete should remove task")
	}
}

func TestDisplayTasksSearch(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{
		{ID: "1", Title: "Buy milk"},
		{ID: "2", Title: "Walk dog"},
	}
	m.SearchActive = true
	m.SearchInput.SetValue("milk")
	got := m.DisplayTasks()
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("expected 1 match for milk, got %d", len(got))
	}
}

func TestToggleCursorSetsCompletedAt(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{{ID: "1", Title: "t", Completed: false}}
	m.ToggleCursorInDisplay(0)
	if !m.Tasks[0].Completed || m.Tasks[0].CompletedAt == nil {
		t.Fatal("expected completed with timestamp")
	}
	m.ToggleCursorInDisplay(0)
	if m.Tasks[0].Completed || m.Tasks[0].CompletedAt != nil {
		t.Fatal("expected incomplete with nil CompletedAt")
	}
}

func TestThemePickerApplyAndCancel(t *testing.T) {
	m := New(config.Config{Theme: "catppuccin"})
	startIdx := m.ThemeIdx
	startID := m.Theme.ID

	m.OpenThemePicker()
	if m.View != ViewTheme {
		t.Fatal("expected ViewTheme")
	}

	// Move to a different theme if possible
	if len(theme.Themes) > 1 {
		m.PreviewThemeDelta(1)
		if m.ThemeIdx == startIdx {
			t.Fatal("preview should change theme index")
		}
	}

	m.CancelThemePicker()
	if m.View != ViewList {
		t.Fatal("expected ViewList after cancel")
	}
	if m.ThemeIdx != startIdx || m.Theme.ID != startID {
		t.Fatalf("cancel should restore theme, got idx=%d id=%q", m.ThemeIdx, m.Theme.ID)
	}
	if m.Config.Theme != "catppuccin" {
		t.Fatalf("cancel should not persist config theme, got %q", m.Config.Theme)
	}

	m.OpenThemePicker()
	m.PreviewTheme(0)
	for i, th := range theme.Themes {
		if th.ID == "nord" {
			m.PreviewTheme(i)
			break
		}
	}
	m.ApplyTheme()
	if m.Config.Theme != "nord" {
		t.Fatalf("apply should persist theme, got %q", m.Config.Theme)
	}
	if m.View != ViewList {
		t.Fatal("expected ViewList after apply")
	}
}

func TestResetFilters(t *testing.T) {
	m := New(config.Config{})
	m.FilterCategory = storage.CategoryPersonal
	m.FilterPriority = storage.PriorityHigh
	m.FilterStatus = FilterStatusPending
	m.SortField = SortTitle
	m.SortAsc = true
	m.ResetFilters()
	if m.HasActiveFilter() {
		t.Fatal("expected no active filter after reset")
	}
	if m.SortField != SortDefault || m.SortAsc {
		t.Fatal("expected default sort after reset")
	}
}

func TestApplyDefaultStatusFilter(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{
		{ID: "1", Title: "Open", Completed: false},
		{ID: "2", Title: "Done", Completed: true},
	}
	m.ApplyDefaultStatusFilter()
	if m.FilterStatus != FilterStatusPending {
		t.Fatalf("expected pending filter, got %v", m.FilterStatus)
	}
	got := m.DisplayTasks()
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("expected only pending task, got %d tasks", len(got))
	}

	m.ResetFilters()
	m.Tasks = []storage.Task{{ID: "2", Title: "Done", Completed: true}}
	m.ApplyDefaultStatusFilter()
	if m.FilterStatus != FilterStatusAll {
		t.Fatalf("expected all when no pending, got %v", m.FilterStatus)
	}
}

func TestClampCursor(t *testing.T) {
	m := New(config.Config{})
	m.Tasks = []storage.Task{{ID: "1", Title: "only"}}
	m.Cursor = 99
	m.ClampCursor()
	if m.Cursor != 0 {
		t.Fatalf("expected cursor 0, got %d", m.Cursor)
	}

	m.Tasks = nil
	m.Cursor = 5
	m.ClampCursor()
	if m.Cursor != 0 {
		t.Fatalf("empty list: expected cursor 0, got %d", m.Cursor)
	}
}
