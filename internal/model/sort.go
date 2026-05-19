package model

import (
	"sort"
	"strings"

	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/storage"
)

// SortField selects how the task list is ordered.
type SortField int

const (
	SortDefault SortField = iota
	SortTitle
	SortPriority
	SortCategory
)

// SortTasks applies the current sort field and direction to a copy of tasks.
func (m *Model) SortTasksSlice(tasks []storage.Task) []storage.Task {
	out := append([]storage.Task(nil), tasks...)
	sort.SliceStable(out, func(i, j int) bool {
		less := m.compareTasks(out[i], out[j])
		if m.SortAsc {
			return less
		}
		return !less
	})
	return out
}

func (m *Model) compareTasks(a, b storage.Task) bool {
	switch m.SortField {
	case SortTitle:
		return strings.ToLower(a.Title) < strings.ToLower(b.Title)
	case SortPriority:
		return storage.PriorityOrder(a.Priority) < storage.PriorityOrder(b.Priority)
	case SortCategory:
		return string(a.Category) < string(b.Category)
	default:
		if a.Completed != b.Completed {
			return !a.Completed
		}
		if m.SortAsc {
			return a.CreatedAt.Before(b.CreatedAt)
		}
		return a.CreatedAt.After(b.CreatedAt)
	}
}

func (m *Model) SortFieldLabel() string {
	switch m.SortField {
	case SortTitle:
		return i18n.T(i18n.KeyTitle)
	case SortPriority:
		return i18n.T(i18n.KeyPriority)
	case SortCategory:
		return i18n.T(i18n.KeyCategory)
	default:
		return i18n.T(i18n.KeySortDefault)
	}
}

func (m *Model) SortDirectionLabel() string {
	if m.SortAsc {
		return i18n.T(i18n.KeySortAsc)
	}
	return i18n.T(i18n.KeySortDesc)
}

func (m *Model) HasCustomSort() bool {
	return m.SortField != SortDefault || m.SortAsc
}
