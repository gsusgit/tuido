package model

import (
	"testing"
	"time"

	"github.com/gsus/todo-app/tui/internal/storage"
)

func TestSortByTitleAsc(t *testing.T) {
	m := Model{SortField: SortTitle, SortAsc: true}
	tasks := []storage.Task{
		{Title: "zebra", CreatedAt: time.Now()},
		{Title: "alpha", CreatedAt: time.Now()},
	}
	out := m.SortTasksSlice(tasks)
	if out[0].Title != "alpha" {
		t.Fatalf("expected alpha first, got %s", out[0].Title)
	}
}

func TestSortByPriorityDesc(t *testing.T) {
	m := Model{SortField: SortPriority, SortAsc: false}
	tasks := []storage.Task{
		{Title: "a", Priority: storage.PriorityLow, CreatedAt: time.Now()},
		{Title: "b", Priority: storage.PriorityHigh, CreatedAt: time.Now()},
	}
	out := m.SortTasksSlice(tasks)
	if out[0].Priority != storage.PriorityLow {
		t.Fatalf("expected low priority first when desc, got %s", out[0].Priority)
	}
}
