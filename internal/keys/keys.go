package keys

import "github.com/charmbracelet/bubbles/key"

// ListKeyMap holds bindings for the list view.
type ListKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	New    key.Binding
	Edit   key.Binding
	Delete key.Binding
	Filter key.Binding
	Search key.Binding
	Toggle key.Binding
	Theme  key.Binding
	Help   key.Binding
	Quit   key.Binding
}

// DefaultListKeyMap is the default list keymap.
var DefaultListKeyMap = ListKeyMap{
	Up:     key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "navigate")),
	Down:   key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "navigate")),
	New:    key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
	Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
	Delete: key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d", "delete")),
	Filter: key.NewBinding(key.WithKeys("f"), key.WithHelp("f", "filter")),
	Search: key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	Toggle: key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "complete")),
	Theme:  key.NewBinding(key.WithKeys("t"), key.WithHelp("t", "theme")),
	Help:   key.NewBinding(key.WithKeys("c", "?", "f1"), key.WithHelp("c", "controls")),
	Quit:   key.NewBinding(key.WithKeys("esc", "ctrl+c", "q"), key.WithHelp("esc", "quit")),
}

// ShortHelp implements help.KeyMap (compact list footer: n, e, d, c).
func (k ListKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.New, k.Edit, k.Delete, k.Help}
}

// FullHelp implements help.KeyMap.
func (k ListKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.New, k.Edit},
		{k.Delete, k.Filter, k.Search, k.Toggle},
		{k.Theme, k.Help, k.Quit},
	}
}

// Matches reports whether msg matches any of the bindings (via String()).
func Matches(keyStr string, bindings ...key.Binding) bool {
	for _, b := range bindings {
		for _, k := range b.Keys() {
			if keyStr == k {
				return true
			}
		}
	}
	return false
}
