package theme

import "github.com/charmbracelet/lipgloss"

// Theme defines a color palette for the TUI.
type Theme struct {
	Name string

	// Base
	Background lipgloss.Color
	Foreground lipgloss.Color

	// UI elements
	Border     lipgloss.Color
	HelpFooter lipgloss.Color
	LogoColor1 lipgloss.Color
	LogoColor2 lipgloss.Color

	// Task states
	TaskTitle    lipgloss.Color
	TaskPending  lipgloss.Color
	TaskDone     lipgloss.Color
	TaskDoneText lipgloss.Color

	// Priority colors
	PriorityHigh   lipgloss.Color
	PriorityMedium lipgloss.Color
	PriorityLow    lipgloss.Color

	// Accent
	Accent  lipgloss.Color
	Accent2 lipgloss.Color

	// Selection
	SelectedBg lipgloss.Color
	SelectedFg lipgloss.Color

	// Modal
	ModalBg  lipgloss.Color
	ModalFg  lipgloss.Color
	Overlay  lipgloss.Color

	// Hotkey
	HotkeyBg lipgloss.Color
	HotkeyFg lipgloss.Color

	// Input
	InputBg lipgloss.Color

	// List
	ListBg lipgloss.Color
}

var (
	CatppuccinMocha = Theme{
		Name: "Catppuccin Mocha",

		Background: lipgloss.Color("#1e1e2e"),
		Foreground: lipgloss.Color("#cdd6f4"),

		Border:     lipgloss.Color("#45475a"),
		HelpFooter: lipgloss.Color("#6c7086"),
		LogoColor1: lipgloss.Color("#cba6f7"),
		LogoColor2: lipgloss.Color("#f5c2e7"),

		TaskTitle:    lipgloss.Color("#cdd6f4"),
		TaskPending:  lipgloss.Color("#a6adc8"),
		TaskDone:     lipgloss.Color("#a6adc8"),
		TaskDoneText: lipgloss.Color("#585b70"),

		PriorityHigh:   lipgloss.Color("#f38ba8"),
		PriorityMedium: lipgloss.Color("#fab387"),
		PriorityLow:    lipgloss.Color("#a6e3a1"),

		Accent:  lipgloss.Color("#89b4fa"),
		Accent2: lipgloss.Color("#f5c2e7"),

		SelectedBg: lipgloss.Color("#45475a"),
		SelectedFg: lipgloss.Color("#cdd6f4"),

		ModalBg: lipgloss.Color("#282837"),
		ModalFg: lipgloss.Color("#cdd6f4"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#45475a"),
		HotkeyFg: lipgloss.Color("#cdd6f4"),

		InputBg: lipgloss.Color("#282839"),

		ListBg: lipgloss.Color("#1c1c28"),
	}

	TokyoNight = Theme{
		Name: "Tokyo Night",

		Background: lipgloss.Color("#1a1b26"),
		Foreground: lipgloss.Color("#a9b1d6"),

		Border:     lipgloss.Color("#32344a"),
		HelpFooter: lipgloss.Color("#565f89"),
		LogoColor1: lipgloss.Color("#7aa2f7"),
		LogoColor2: lipgloss.Color("#bb9af7"),

		TaskTitle:    lipgloss.Color("#a9b1d6"),
		TaskPending:  lipgloss.Color("#9aa5ce"),
		TaskDone:     lipgloss.Color("#565f89"),
		TaskDoneText: lipgloss.Color("#3b4261"),

		PriorityHigh:   lipgloss.Color("#f7768e"),
		PriorityMedium: lipgloss.Color("#e0af68"),
		PriorityLow:    lipgloss.Color("#9ece6a"),

		Accent:  lipgloss.Color("#7aa2f7"),
		Accent2: lipgloss.Color("#bb9af7"),

		SelectedBg: lipgloss.Color("#2f3a55"),
		SelectedFg: lipgloss.Color("#a9b1d6"),

		ModalBg: lipgloss.Color("#282837"),
		ModalFg: lipgloss.Color("#a9b1d6"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#2f3a55"),
		HotkeyFg: lipgloss.Color("#a9b1d6"),

		InputBg: lipgloss.Color("#1f2030"),

		ListBg: lipgloss.Color("#181920"),
	}

	OneDark = Theme{
		Name: "One Dark",

		Background: lipgloss.Color("#282c34"),
		Foreground: lipgloss.Color("#abb2bf"),

		Border:     lipgloss.Color("#3e4452"),
		HelpFooter: lipgloss.Color("#5c6370"),
		LogoColor1: lipgloss.Color("#61afef"),
		LogoColor2: lipgloss.Color("#e06c75"),

		TaskTitle:    lipgloss.Color("#abb2bf"),
		TaskPending:  lipgloss.Color("#828997"),
		TaskDone:     lipgloss.Color("#5c6370"),
		TaskDoneText: lipgloss.Color("#3e4452"),

		PriorityHigh:   lipgloss.Color("#e06c75"),
		PriorityMedium: lipgloss.Color("#d19a66"),
		PriorityLow:    lipgloss.Color("#98c379"),

		Accent:  lipgloss.Color("#61afef"),
		Accent2: lipgloss.Color("#c678dd"),

		SelectedBg: lipgloss.Color("#3e4452"),
		SelectedFg: lipgloss.Color("#abb2bf"),

		ModalBg: lipgloss.Color("#282837"),
		ModalFg: lipgloss.Color("#abb2bf"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#3e4452"),
		HotkeyFg: lipgloss.Color("#abb2bf"),

		InputBg: lipgloss.Color("#2a2e36"),

		ListBg: lipgloss.Color("#262830"),
	}

	Monochrome = Theme{
		Name: "Monochrome",

		Background: lipgloss.Color("#000000"),
		Foreground: lipgloss.Color("#ffffff"),

		Border:     lipgloss.Color("#888888"),
		HelpFooter: lipgloss.Color("#666666"),
		LogoColor1: lipgloss.Color("#ffffff"),
		LogoColor2: lipgloss.Color("#aaaaaa"),

		TaskTitle:    lipgloss.Color("#ffffff"),
		TaskPending:  lipgloss.Color("#cccccc"),
		TaskDone:     lipgloss.Color("#555555"),
		TaskDoneText: lipgloss.Color("#444444"),

		PriorityHigh:   lipgloss.Color("#ffffff"),
		PriorityMedium: lipgloss.Color("#aaaaaa"),
		PriorityLow:    lipgloss.Color("#555555"),

		Accent:  lipgloss.Color("#ffffff"),
		Accent2: lipgloss.Color("#888888"),

		SelectedBg: lipgloss.Color("#333333"),
		SelectedFg: lipgloss.Color("#ffffff"),

		ModalBg: lipgloss.Color("#282837"),
		ModalFg: lipgloss.Color("#ffffff"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#333333"),
		HotkeyFg: lipgloss.Color("#ffffff"),

		InputBg: lipgloss.Color("#1a1a1a"),

		ListBg: lipgloss.Color("#181818"),
	}

	Themes = []Theme{CatppuccinMocha, TokyoNight, OneDark, Monochrome}
)

// ThemeID returns the config id for a theme.
func ThemeID(t Theme) string {
	switch t.Name {
	case TokyoNight.Name:
		return "tokyo-night"
	case OneDark.Name:
		return "one-dark"
	case Monochrome.Name:
		return "monochrome"
	default:
		return "catppuccin"
	}
}

// ByID returns theme for config id.
func ByID(id string) (Theme, int) {
	switch id {
	case "tokyo-night":
		return TokyoNight, 1
	case "one-dark":
		return OneDark, 2
	case "monochrome":
		return Monochrome, 3
	default:
		return CatppuccinMocha, 0
	}
}