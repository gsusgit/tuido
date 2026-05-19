package model

import (
	"testing"

	"github.com/gsus/todo-app/tui/internal/config"
	"github.com/gsus/todo-app/tui/internal/storage"
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
