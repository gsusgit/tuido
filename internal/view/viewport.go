package view

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gsusgit/tuido/internal/i18n"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/theme"
)

// Minimum terminal size required to render the main UI (columns × lines).
const (
	MinWidth  = 60
	MinHeight = 20
)

// TerminalSize describes the current terminal dimensions.
type TerminalSize struct {
	Width  int
	Height int
}

// SizeFromModel reads dimensions from the application model.
func SizeFromModel(m *model.Model) TerminalSize {
	return TerminalSize{Width: m.Width, Height: m.Height}
}

// IsMeasurable reports whether Bubble Tea has reported a real size yet.
func (s TerminalSize) IsMeasurable() bool {
	return s.Width > 0 && s.Height > 0
}

// MeetsMinimum reports whether dimensions satisfy MinWidth×MinHeight.
func (s TerminalSize) MeetsMinimum() bool {
	return s.Width >= MinWidth && s.Height >= MinHeight
}

// IsValid is true when the main UI can be rendered safely.
func (s TerminalSize) IsValid() bool {
	return s.IsMeasurable() && s.MeetsMinimum()
}

// RenderViewport gates rendering behind a minimum size check.
// Pass a callback that builds the main UI for the current view.
//
// Example for a future panel:
//
//	func renderMyPanel(m *model.Model) string {
//	    return view.RenderViewport(m, func(m *model.Model) string {
//	        return buildMyPanelContent(m)
//	    })
//	}
func RenderViewport(m *model.Model, renderMain func(*model.Model) string) string {
	size := SizeFromModel(m)
	if !size.IsValid() {
		return RenderTerminalTooSmall(m.Theme, size)
	}
	return renderMain(m)
}

// RenderTerminalTooSmall draws a centered fallback when the terminal is too small.
func RenderTerminalTooSmall(t theme.Theme, size TerminalSize) string {
	w := size.Width
	h := size.Height
	if w < 1 {
		w = MinWidth
	}
	if h < 1 {
		h = MinHeight
	}

	title := lipgloss.NewStyle().
		Foreground(t.Accent).
		Bold(true).
		Render(i18n.T(i18n.KeyWindowTooSmall))

	subtitle := lipgloss.NewStyle().
		Foreground(t.HelpFooter).
		Render(i18n.T(i18n.KeyMinimumSizeRequired, MinWidth, MinHeight))

	block := lipgloss.JoinVertical(lipgloss.Center, title, subtitle)
	return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, block)
}
