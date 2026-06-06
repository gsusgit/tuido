package tuiapp

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gsusgit/tuido/internal/config"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/keys"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/storage"
	"github.com/gsusgit/tuido/internal/view"
)

const toastDuration = 2 * time.Second

// Program is the Bubble Tea application model.
type Program struct {
	model model.Model
}

func toastDismissCmd() tea.Cmd {
	return tea.Tick(toastDuration, func(time.Time) tea.Msg { return toastTickMsg{} })
}

func (p *Program) flashToast(msg string) tea.Cmd {
	p.model.SetToast(model.ToastInfo, msg)
	return toastDismissCmd()
}

// NewProgram creates the TUI program with loaded tasks.
func NewProgram(cfg config.Config, tasks []storage.Task) *Program {
	m := model.New(cfg)
	m.Tasks = tasks
	m.ApplyDefaultStatusFilter()
	m.SortTasks()
	m.ClampCursor()
	m.BlurInputs()
	return &Program{model: m}
}

func (p *Program) Init() tea.Cmd {
	return p.model.StatsSpinner.Tick
}

func (p *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.model.Width = msg.Width
		p.model.Height = msg.Height
		return p, nil

	case toastTickMsg:
		p.model.ClearToast()
		return p, nil

	case deleteExpireMsg:
		p.model.CancelDeletePending()
		return p, nil

	case saveErrMsg:
		p.model.SetToast(model.ToastError, i18n.T(i18n.KeySaveError, msg.err.Error()))
		return p, toastDismissCmd()

	case saveDoneMsg:
		return p, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		p.model.StatsSpinner, cmd = p.model.StatsSpinner.Update(msg)
		return p, cmd

	case tea.KeyMsg:
		return p.handleKeyMsg(msg)
	}

	return p, nil
}

func (p *Program) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	s := msg.String()

	if p.model.View == model.ViewHelp {
		if s == "esc" || s == "?" || s == "c" || s == "f1" {
			p.model.View = model.ViewList
			p.model.BlurInputs()
			return p, nil
		}
		if s == "ctrl+c" || s == "q" {
			return p, p.SaveAndQuit()
		}
		return p, nil
	}

	// Quit
	if p.shouldQuit(s) {
		return p, p.SaveAndQuit()
	}

	var cmds []tea.Cmd

	if p.model.View == model.ViewList && p.model.SearchActive {
		switch s {
		case "esc":
			p.model.SearchActive = false
			p.model.SearchInput.SetValue("")
			p.model.SearchInput.Blur()
			return p, nil
		case "enter":
			p.model.SearchActive = false
			p.model.SearchInput.Blur()
			return p, nil
		default:
			var cmd tea.Cmd
			p.model.SearchInput, cmd = p.model.SearchInput.Update(msg)
			return p, cmd
		}
	}

	changed, extra := p.routeKey(s, msg)
	cmds = append(cmds, extra...)

	if changed {
		cmds = append(cmds, p.persistCmd())
	}

	if len(cmds) == 0 {
		return p, nil
	}
	return p, tea.Batch(cmds...)
}

func (p *Program) shouldQuit(s string) bool {
	switch p.model.View {
	case model.ViewList:
		if p.model.SearchActive {
			return false
		}
		if p.model.DeletePendingID != "" && s == "esc" {
			p.model.CancelDeletePending()
			return false
		}
		return s == "ctrl+c" || s == "q" || s == "esc"
	case model.ViewInputTask, model.ViewFilter, model.ViewTheme:
		return s == "ctrl+c"
	}
	return false
}

func (p *Program) routeKey(s string, msg tea.KeyMsg) (changed bool, extra []tea.Cmd) {
	km := p.model.KeyMap

	switch p.model.View {
	case model.ViewList:
		p.model.BlurInputs()

		if p.model.DeletePendingID != "" {
			switch s {
			case "enter":
				_, deleted := p.model.DeleteByDisplayIndex(p.model.Cursor)
				if deleted {
					extra = append(extra, p.flashToast(i18n.T(i18n.KeyTaskDeleted)))
					return true, extra
				}
				return false, extra
			case "esc":
				p.model.CancelDeletePending()
				return false, nil
			case "d", "x":
				return false, nil
			}
		}

		switch s {
		case "up", "k":
			p.model.MoveCursor(-1)
			return false, nil
		case "down", "j":
			p.model.MoveCursor(1)
			return false, nil
		case " ":
			p.model.ToggleCursorInDisplay(p.model.Cursor)
			return true, nil
		case "?", "c", "f1":
			p.model.View = model.ViewHelp
			p.model.BlurInputs()
			return false, nil
		case "n":
			cmd := p.openNewTask()
			return false, []tea.Cmd{cmd}
		case "e":
			if len(p.model.DisplayTasks()) > 0 {
				cmd := p.model.LoadTaskForEdit(p.model.Cursor)
				return false, []tea.Cmd{cmd}
			}
			return false, nil
		case "f":
			p.model.OpenFilter()
			return false, nil
		case "r":
			if p.model.HasActiveFilter() {
				p.model.ResetFilters()
			}
			return false, nil
		case "d", "x":
			pending, deleted := p.model.DeleteByDisplayIndex(p.model.Cursor)
			if deleted {
				extra = append(extra, p.flashToast(i18n.T(i18n.KeyTaskDeleted)))
				return true, extra
			}
			if pending {
				extra = append(extra, tea.Tick(3*time.Second, func(time.Time) tea.Msg { return deleteExpireMsg{} }))
			}
			return false, extra
		case "t":
			p.model.OpenThemePicker()
			return false, nil
		case "/":
			p.model.SearchActive = true
			p.model.SearchInput.Focus()
			p.model.SearchInput.SetValue("")
			return false, nil
		}

		if keys.Matches(s, km.Up) {
			p.model.MoveCursor(-1)
		}
		if keys.Matches(s, km.Down) {
			p.model.MoveCursor(1)
		}

	case model.ViewInputTask:
		switch s {
		case "esc":
			p.model.ClearInput()
			p.model.View = model.ViewList
			p.model.BlurInputs()
			return false, nil
		case "tab":
			p.model.InputCursor = (p.model.InputCursor + 1) % 3
			if p.model.InputCursor == 0 {
				cmd := p.model.TitleInput.Focus()
				return false, []tea.Cmd{cmd}
			} else {
				p.model.TitleInput.Blur()
			}
			return false, nil
		case "enter":
			title := p.model.TitleInput.Value()
			if title == "" {
				p.model.InputErr = i18n.T(i18n.KeyTitleRequired)
				p.model.InputCursor = 0
				p.model.TitleInput.Focus()
				return false, nil
			}
			var toastCmd tea.Cmd
			if p.model.EditingTaskID != "" {
				p.model.SaveEditedTask(title, p.model.InputCategory, p.model.InputPriority)
				toastCmd = p.flashToast(i18n.T(i18n.KeyTaskSaved))
			} else if err := p.model.AddTask(title, p.model.InputCategory, p.model.InputPriority); err != nil {
				p.model.InputErr = err.Error()
				return false, nil
			} else {
				toastCmd = p.flashToast(i18n.T(i18n.KeyTaskCreated))
			}
			p.model.View = model.ViewList
			p.model.BlurInputs()
			return true, []tea.Cmd{toastCmd}
		case "left", "right":
			if p.model.InputCursor == 0 {
				var cmd tea.Cmd
				p.model.TitleInput, cmd = p.model.TitleInput.Update(msg)
				return false, []tea.Cmd{cmd}
			}
			p.cycleInputOption(s == "left")
			return false, nil
		default:
			if p.model.InputCursor == 0 {
				var cmd tea.Cmd
				p.model.TitleInput, cmd = p.model.TitleInput.Update(msg)
				return false, []tea.Cmd{cmd}
			}
		}

	case model.ViewFilter:
		switch s {
		case "esc":
			p.model.View = model.ViewList
			p.model.BlurInputs()
			return false, nil
		case "r":
			p.model.ResetFilters()
			return false, nil
		case "enter":
			p.model.Cursor = 0
			p.model.View = model.ViewList
			p.model.BlurInputs()
			return false, nil
		case "tab":
			p.model.FilterCursor = (p.model.FilterCursor + 1) % 5
			return false, nil
		case "left", "right":
			p.cycleFilterOption(s == "left")
			return false, nil
		}

	case model.ViewTheme:
		switch s {
		case "esc":
			p.model.CancelThemePicker()
			return false, nil
		case "enter":
			p.model.ApplyTheme()
			return true, nil
		case "up", "k":
			p.model.PreviewThemeDelta(-1)
			return false, nil
		case "down", "j":
			p.model.PreviewThemeDelta(1)
			return false, nil
		}
	}

	return false, nil
}

func (p *Program) openNewTask() tea.Cmd {
	p.model.ClearInput()
	p.model.View = model.ViewInputTask
	p.model.InputCursor = 0
	return p.model.TitleInput.Focus()
}

func (p *Program) cycleInputOption(left bool) {
	if p.model.InputCursor == 1 {
		cats := storage.AllCategories()
		idx := indexOfCategory(p.model.InputCategory, cats)
		if left {
			idx = (idx + len(cats) - 1) % len(cats)
		} else {
			idx = (idx + 1) % len(cats)
		}
		p.model.InputCategory = cats[idx]
	} else if p.model.InputCursor == 2 {
		prios := storage.AllPriorities()
		idx := indexOfPriority(p.model.InputPriority, prios)
		if left {
			idx = (idx + len(prios) - 1) % len(prios)
		} else {
			idx = (idx + 1) % len(prios)
		}
		p.model.InputPriority = prios[idx]
	}
}

func (p *Program) cycleFilterOption(left bool) {
	switch p.model.FilterCursor {
	case 0:
		vals := []storage.Category{"", storage.CategoryPersonal, storage.CategoryTrabajo}
		idx := indexOfCategory(p.model.FilterCategory, vals)
		if left {
			idx = (idx + len(vals) - 1) % len(vals)
		} else {
			idx = (idx + 1) % len(vals)
		}
		p.model.FilterCategory = vals[idx]
	case 1:
		vals := []storage.Priority{"", storage.PriorityHigh, storage.PriorityMedium, storage.PriorityLow}
		idx := indexOfPriority(p.model.FilterPriority, vals)
		if left {
			idx = (idx + len(vals) - 1) % len(vals)
		} else {
			idx = (idx + 1) % len(vals)
		}
		p.model.FilterPriority = vals[idx]
	case 2:
		vals := []model.FilterStatus{
			model.FilterStatusAll, model.FilterStatusPending, model.FilterStatusCompleted,
		}
		idx := indexOfFilterStatus(p.model.FilterStatus, vals)
		if left {
			idx = (idx + len(vals) - 1) % len(vals)
		} else {
			idx = (idx + 1) % len(vals)
		}
		p.model.FilterStatus = vals[idx]
	case 3:
		fields := []model.SortField{
			model.SortDefault, model.SortTitle, model.SortPriority, model.SortCategory,
		}
		idx := indexOfSortField(p.model.SortField, fields)
		if left {
			idx = (idx + len(fields) - 1) % len(fields)
		} else {
			idx = (idx + 1) % len(fields)
		}
		p.model.SortField = fields[idx]
	case 4:
		// UI order: Desc ↓ (left) · Asc ↑ (right)
		if left {
			p.model.SortAsc = false
		} else {
			p.model.SortAsc = true
		}
	}
}

func indexOfFilterStatus(s model.FilterStatus, list []model.FilterStatus) int {
	for i, v := range list {
		if v == s {
			return i
		}
	}
	return 0
}

func indexOfSortField(s model.SortField, list []model.SortField) int {
	for i, v := range list {
		if v == s {
			return i
		}
	}
	return 0
}

func indexOfCategory(c storage.Category, list []storage.Category) int {
	for i, v := range list {
		if v == c {
			return i
		}
	}
	return 0
}

func indexOfPriority(pr storage.Priority, list []storage.Priority) int {
	for i, v := range list {
		if v == pr {
			return i
		}
	}
	return 0
}

func (p *Program) persistCmd() tea.Cmd {
	return func() tea.Msg {
		if err := storage.Save(p.model.Tasks); err != nil {
			return saveErrMsg{err: err}
		}
		_ = config.Save(p.model.Config)
		return saveDoneMsg{}
	}
}

func (p *Program) SaveAndQuit() tea.Cmd {
	return tea.Sequence(
		func() tea.Msg {
			if err := storage.Save(p.model.Tasks); err != nil {
				return saveErrMsg{err: err}
			}
			_ = config.Save(p.model.Config)
			return saveDoneMsg{}
		},
		tea.Quit,
	)
}

func (p *Program) View() string {
	return view.Render(&p.model)
}

// Model returns inner state for testing.
func (p *Program) Model() model.Model { return p.model }
