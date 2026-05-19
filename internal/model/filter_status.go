package model

import "github.com/gsus/todo-app/tui/internal/i18n"

// FilterStatus limits the list by task completion.
type FilterStatus int

const (
	FilterStatusAll FilterStatus = iota
	FilterStatusPending
	FilterStatusCompleted
)

func (m *Model) CompletionStatusLabel(s FilterStatus) string {
	switch s {
	case FilterStatusPending:
		return i18n.T(i18n.KeyPending)
	case FilterStatusCompleted:
		return i18n.T(i18n.KeyStatusCompleted)
	default:
		return i18n.T(i18n.KeyAll)
	}
}
