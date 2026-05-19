package model

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/config"
	"github.com/gsus/todo-app/tui/internal/i18n"
	"github.com/gsus/todo-app/tui/internal/keys"
	"github.com/gsus/todo-app/tui/internal/storage"
	"github.com/gsus/todo-app/tui/internal/theme"
)

type View int

const (
	ViewList View = iota
	ViewInputTask
	ViewFilter
	ViewHelp
)

// ToastKind classifies flash messages.
type ToastKind int

const (
	ToastNone ToastKind = iota
	ToastInfo
	ToastError
)

type Model struct {
	Tasks    []storage.Task
	Cursor   int
	View     View
	ThemeIdx int
	Theme    theme.Theme
	Config   config.Config

	TitleInput  textinput.Model
	SearchInput textinput.Model
	InputCursor int
	InputCategory storage.Category
	InputPriority storage.Priority
	InputErr      string
	EditingTaskID string

	FilterCategory storage.Category
	FilterPriority storage.Priority
	FilterStatus   FilterStatus
	FilterCursor   int
	SortField      SortField
	SortAsc        bool

	SearchActive bool

	DeletePendingID string

	Toast     string
	ToastKind ToastKind

	ListVP       viewport.Model
	Help         help.Model
	StatsSpinner spinner.Model
	KeyMap       keys.ListKeyMap

	Width  int
	Height int

	// ListBodyH is the viewport height in rows (set on resize).
	ListBodyH int
}

// New builds a model from persisted config.
func New(cfg config.Config) Model {
	th, idx := theme.ByID(cfg.Theme)
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = "..."
	ti.CharLimit = 200
	ti.Width = 40

	si := textinput.New()
	si.Prompt = ""
	si.Placeholder = i18n.T(i18n.KeySearch)
	si.CharLimit = 100
	si.Width = 30

	h := help.New()
	km := keys.DefaultListKeyMap

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(th.Accent)

	return Model{
		Tasks:         []storage.Task{},
		View:          ViewList,
		ThemeIdx:      idx,
		Theme:         th,
		Config:        cfg,
		TitleInput:    ti,
		SearchInput:   si,
		InputCategory: storage.CategoryPersonal,
		InputPriority: storage.PriorityMedium,
		ListVP:        viewport.New(80, 20),
		Help:          h,
		StatsSpinner:  sp,
		KeyMap:        km,
	}
}

func (m *Model) HasActiveFilter() bool {
	return m.FilterCategory != "" || m.FilterPriority != "" || m.FilterStatus != FilterStatusAll || m.HasCustomSort()
}

func (m *Model) DisplayTasks() []storage.Task {
	tasks := m.Tasks
	if m.HasActiveFilter() {
		tasks = m.FilteredTasks()
	}
	if m.SearchActive && m.SearchInput.Value() != "" {
		q := strings.ToLower(m.SearchInput.Value())
		var out []storage.Task
		for _, t := range tasks {
			if strings.Contains(strings.ToLower(t.Title), q) {
				out = append(out, t)
			}
		}
		tasks = out
	}
	return m.SortTasksSlice(tasks)
}

func (m *Model) FilteredTasks() []storage.Task {
	var result []storage.Task
	for _, t := range m.Tasks {
		if m.FilterCategory != "" && t.Category != m.FilterCategory {
			continue
		}
		if m.FilterPriority != "" && t.Priority != m.FilterPriority {
			continue
		}
		switch m.FilterStatus {
		case FilterStatusPending:
			if t.Completed {
				continue
			}
		case FilterStatusCompleted:
			if !t.Completed {
				continue
			}
		}
		result = append(result, t)
	}
	return result
}

func (m *Model) NextTheme() {
	m.ThemeIdx = (m.ThemeIdx + 1) % len(theme.Themes)
	m.Theme = theme.Themes[m.ThemeIdx]
	m.Config.Theme = theme.ThemeID(m.Theme)
	m.StatsSpinner.Style = lipgloss.NewStyle().Foreground(m.Theme.Accent)
}

func (m *Model) SortTasks() {
	storage.SortTasks(m.Tasks)
}

func (m *Model) AddTask(title string, cat storage.Category, prio storage.Priority) error {
	id, err := storage.NewID()
	if err != nil {
		return err
	}
	now := time.Now()
	task := storage.Task{
		ID:        id,
		Title:     title,
		Completed: false,
		Priority:  prio,
		Category:  cat,
		CreatedAt: now,
	}
	m.Tasks = append([]storage.Task{task}, m.Tasks...)
	m.SortTasks()
	m.Cursor = 0
	m.SyncViewportOffset()
	return nil
}

func (m *Model) LoadTaskForEdit(displayIdx int) {
	display := m.DisplayTasks()
	if displayIdx < 0 || displayIdx >= len(display) {
		return
	}
	task := display[displayIdx]
	m.Cursor = displayIdx
	m.EditingTaskID = task.ID
	m.TitleInput.SetValue(task.Title)
	m.InputCategory = task.Category
	m.InputPriority = task.Priority
	m.InputCursor = 0
	m.InputErr = ""
	m.View = ViewInputTask
}

func (m *Model) SaveEditedTask(title string, cat storage.Category, prio storage.Priority) {
	if m.EditingTaskID == "" {
		return
	}
	for i := range m.Tasks {
		if m.Tasks[i].ID == m.EditingTaskID {
			m.Tasks[i].Title = title
			m.Tasks[i].Category = cat
			m.Tasks[i].Priority = prio
			break
		}
	}
	m.SortTasks()
	m.ClearInput()
}

func (m *Model) ClearInput() {
	m.EditingTaskID = ""
	m.TitleInput.SetValue("")
	m.TitleInput.Blur()
	m.InputCategory = storage.CategoryPersonal
	m.InputPriority = storage.PriorityMedium
	m.InputCursor = 0
	m.InputErr = ""
}

func (m *Model) BlurInputs() {
	m.TitleInput.Blur()
	m.SearchInput.Blur()
}

func (m *Model) OpenFilter() {
	m.FilterCursor = 0
	m.View = ViewFilter
}

func (m *Model) ResetFilters() {
	m.FilterCategory = ""
	m.FilterPriority = ""
	m.FilterStatus = FilterStatusAll
	m.SortField = SortDefault
	m.SortAsc = false
	m.Cursor = 0
	m.ClampCursor()
	m.SyncViewportOffset()
}

func (m *Model) ToggleCursorInDisplay(displayIdx int) {
	display := m.DisplayTasks()
	if displayIdx < 0 || displayIdx >= len(display) {
		return
	}
	id := display[displayIdx].ID
	for i := range m.Tasks {
		if m.Tasks[i].ID == id {
			t := &m.Tasks[i]
			t.Completed = !t.Completed
			if t.Completed {
				now := time.Now()
				t.CompletedAt = &now
			} else {
				t.CompletedAt = nil
			}
			m.SortTasks()
			m.ClampCursor()
			return
		}
	}
}

func (m *Model) DeleteByDisplayIndex(displayIdx int) (pending bool, deleted bool) {
	display := m.DisplayTasks()
	if displayIdx < 0 || displayIdx >= len(display) {
		return false, false
	}
	id := display[displayIdx].ID
	if m.DeletePendingID == id {
		m.deleteByID(id)
		m.DeletePendingID = ""
		return false, true
	}
	m.DeletePendingID = id
	return true, false
}

func (m *Model) CancelDeletePending() {
	m.DeletePendingID = ""
}

func (m *Model) deleteByID(id string) {
	for i, t := range m.Tasks {
		if t.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			m.ClampCursor()
			m.SyncViewportOffset()
			return
		}
	}
}

func (m *Model) ClampCursor() {
	display := m.DisplayTasks()
	if len(display) == 0 {
		m.Cursor = 0
		return
	}
	if m.Cursor >= len(display) {
		m.Cursor = len(display) - 1
	}
	if m.Cursor < 0 {
		m.Cursor = 0
	}
}

func (m *Model) MoveCursor(delta int) {
	display := m.DisplayTasks()
	if len(display) == 0 {
		return
	}
	m.Cursor = (m.Cursor + delta + len(display)) % len(display)
	m.SyncViewportOffset()
}

// SyncViewportOffset keeps the cursor row visible in the list viewport.
func (m *Model) SyncViewportOffset() {
	display := m.DisplayTasks()
	if len(display) == 0 {
		m.ListVP.YOffset = 0
		return
	}
	if m.Cursor < m.ListVP.YOffset {
		m.ListVP.YOffset = m.Cursor
	}
	if m.Cursor >= m.ListVP.YOffset+m.ListVP.Height {
		m.ListVP.YOffset = m.Cursor - m.ListVP.Height + 1
	}
	if m.ListVP.YOffset < 0 {
		m.ListVP.YOffset = 0
	}
	max := len(display) - m.ListVP.Height
	if max < 0 {
		max = 0
	}
	if m.ListVP.YOffset > max {
		m.ListVP.YOffset = max
	}
}

func (m *Model) Stats() (total, completed int) {
	total = len(m.Tasks)
	for _, t := range m.Tasks {
		if t.Completed {
			completed++
		}
	}
	return
}

func (m *Model) SetToast(kind ToastKind, msg string) {
	m.Toast = msg
	m.ToastKind = kind
}

func (m *Model) ClearToast() {
	m.Toast = ""
	m.ToastKind = ToastNone
}

func (m *Model) CategoryLabel(c storage.Category) string {
	switch c {
	case storage.CategoryTrabajo:
		return i18n.T(i18n.KeyTrabajo)
	default:
		return i18n.T(i18n.KeyPersonal)
	}
}

func (m *Model) PriorityLabel(p storage.Priority) string {
	switch p {
	case storage.PriorityHigh:
		return i18n.T(i18n.KeyPriorityHigh)
	case storage.PriorityLow:
		return i18n.T(i18n.KeyPriorityLow)
	default:
		return i18n.T(i18n.KeyPriorityMedium)
	}
}
