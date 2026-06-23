package theme

import "github.com/charmbracelet/lipgloss"

// Theme defines a color palette for the TUI.
type Theme struct {
	ID   string
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
	ModalBg lipgloss.Color
	ModalFg lipgloss.Color
	Overlay lipgloss.Color

	// Hotkey
	HotkeyBg lipgloss.Color
	HotkeyFg lipgloss.Color

	// Input
	InputBg lipgloss.Color

	// List
	ListBg lipgloss.Color
}

var (
	Catppuccin = Theme{
		ID:   "catppuccin",
		Name: "Catppuccin",

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
		ID:   "tokyo-night",
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
		ID:   "one-dark",
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
		ID:   "monochrome",
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

	Nord = Theme{
		ID:   "nord",
		Name: "Nord",

		Background: lipgloss.Color("#2E3440"),
		Foreground: lipgloss.Color("#D8DEE9"),

		Border:     lipgloss.Color("#4C566A"),
		HelpFooter: lipgloss.Color("#616E88"),
		LogoColor1: lipgloss.Color("#88C0D0"),
		LogoColor2: lipgloss.Color("#B48EAD"),

		TaskTitle:    lipgloss.Color("#D8DEE9"),
		TaskPending:  lipgloss.Color("#E5E9F0"),
		TaskDone:     lipgloss.Color("#616E88"),
		TaskDoneText: lipgloss.Color("#4C566A"),

		PriorityHigh:   lipgloss.Color("#BF616A"),
		PriorityMedium: lipgloss.Color("#EBCB8B"),
		PriorityLow:    lipgloss.Color("#A3BE8C"),

		Accent:  lipgloss.Color("#88C0D0"),
		Accent2: lipgloss.Color("#81A1C1"),

		SelectedBg: lipgloss.Color("#434C5E"),
		SelectedFg: lipgloss.Color("#ECEFF4"),

		ModalBg: lipgloss.Color("#3B4252"),
		ModalFg: lipgloss.Color("#ECEFF4"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#434C5E"),
		HotkeyFg: lipgloss.Color("#ECEFF4"),

		InputBg: lipgloss.Color("#3B4252"),

		ListBg: lipgloss.Color("#2E3440"),
	}

	Gruvbox = Theme{
		ID:   "gruvbox",
		Name: "Gruvbox Dark",

		Background: lipgloss.Color("#282828"),
		Foreground: lipgloss.Color("#EBDBB2"),

		Border:     lipgloss.Color("#504945"),
		HelpFooter: lipgloss.Color("#928374"),
		LogoColor1: lipgloss.Color("#83A598"),
		LogoColor2: lipgloss.Color("#D3869B"),

		TaskTitle:    lipgloss.Color("#EBDBB2"),
		TaskPending:  lipgloss.Color("#D5C4A1"),
		TaskDone:     lipgloss.Color("#928374"),
		TaskDoneText: lipgloss.Color("#665C54"),

		PriorityHigh:   lipgloss.Color("#FB4934"),
		PriorityMedium: lipgloss.Color("#FABD2F"),
		PriorityLow:    lipgloss.Color("#B8BB26"),

		Accent:  lipgloss.Color("#83A598"),
		Accent2: lipgloss.Color("#D3869B"),

		SelectedBg: lipgloss.Color("#504945"),
		SelectedFg: lipgloss.Color("#FBF1C7"),

		ModalBg: lipgloss.Color("#3C3836"),
		ModalFg: lipgloss.Color("#FBF1C7"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#504945"),
		HotkeyFg: lipgloss.Color("#FBF1C7"),

		InputBg: lipgloss.Color("#32302F"),

		ListBg: lipgloss.Color("#1D2021"),
	}

	// Ristretto — Monokai Pro Ristretto (Omarchy); warm browns, not Catppuccin.
	Ristretto = Theme{
		ID:   "ristretto",
		Name: "Ristretto",

		Background: lipgloss.Color("#2c2525"),
		Foreground: lipgloss.Color("#e6d9db"),

		Border:     lipgloss.Color("#5b4a45"),
		HelpFooter: lipgloss.Color("#948a8b"),
		LogoColor1: lipgloss.Color("#f38d70"),
		LogoColor2: lipgloss.Color("#fd6883"),

		TaskTitle:    lipgloss.Color("#e6d9db"),
		TaskPending:  lipgloss.Color("#c3b7b8"),
		TaskDone:     lipgloss.Color("#72696a"),
		TaskDoneText: lipgloss.Color("#5b4a45"),

		PriorityHigh:   lipgloss.Color("#fd6883"),
		PriorityMedium: lipgloss.Color("#f9cc6c"),
		PriorityLow:    lipgloss.Color("#adda78"),

		Accent:  lipgloss.Color("#f38d70"),
		Accent2: lipgloss.Color("#a8a9eb"),

		SelectedBg: lipgloss.Color("#403e41"),
		SelectedFg: lipgloss.Color("#e6d9db"),

		ModalBg: lipgloss.Color("#3d2f2a"),
		ModalFg: lipgloss.Color("#e6d9db"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#403e41"),
		HotkeyFg: lipgloss.Color("#e6d9db"),

		InputBg: lipgloss.Color("#2c2421"),

		ListBg: lipgloss.Color("#241e1e"),
	}

	Monokai = Theme{
		ID:   "monokai",
		Name: "Monokai",

		Background: lipgloss.Color("#272822"),
		Foreground: lipgloss.Color("#F8F8F2"),

		Border:     lipgloss.Color("#49483E"),
		HelpFooter: lipgloss.Color("#75715E"),
		LogoColor1: lipgloss.Color("#66D9EF"),
		LogoColor2: lipgloss.Color("#F92672"),

		TaskTitle:    lipgloss.Color("#F8F8F2"),
		TaskPending:  lipgloss.Color("#E6DB74"),
		TaskDone:     lipgloss.Color("#75715E"),
		TaskDoneText: lipgloss.Color("#49483E"),

		PriorityHigh:   lipgloss.Color("#F92672"),
		PriorityMedium: lipgloss.Color("#FD971F"),
		PriorityLow:    lipgloss.Color("#A6E22E"),

		Accent:  lipgloss.Color("#66D9EF"),
		Accent2: lipgloss.Color("#AE81FF"),

		SelectedBg: lipgloss.Color("#49483E"),
		SelectedFg: lipgloss.Color("#F8F8F2"),

		ModalBg: lipgloss.Color("#3E3D32"),
		ModalFg: lipgloss.Color("#F8F8F2"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#49483E"),
		HotkeyFg: lipgloss.Color("#F8F8F2"),

		InputBg: lipgloss.Color("#2F2F2A"),

		ListBg: lipgloss.Color("#1E1F1C"),
	}

	Darcula = Theme{
		ID:   "darcula",
		Name: "Darcula",

		Background: lipgloss.Color("#2B2B2B"),
		Foreground: lipgloss.Color("#A9B7C6"),

		Border:     lipgloss.Color("#323232"),
		HelpFooter: lipgloss.Color("#808080"),
		LogoColor1: lipgloss.Color("#6897BB"),
		LogoColor2: lipgloss.Color("#9876AA"),

		TaskTitle:    lipgloss.Color("#A9B7C6"),
		TaskPending:  lipgloss.Color("#BBB529"),
		TaskDone:     lipgloss.Color("#606366"),
		TaskDoneText: lipgloss.Color("#4B4B4B"),

		PriorityHigh:   lipgloss.Color("#FF6B68"),
		PriorityMedium: lipgloss.Color("#CC7832"),
		PriorityLow:    lipgloss.Color("#6A8759"),

		Accent:  lipgloss.Color("#6897BB"),
		Accent2: lipgloss.Color("#9876AA"),

		SelectedBg: lipgloss.Color("#214283"),
		SelectedFg: lipgloss.Color("#FFFFFF"),

		ModalBg: lipgloss.Color("#3C3F41"),
		ModalFg: lipgloss.Color("#A9B7C6"),
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color("#3C3F41"),
		HotkeyFg: lipgloss.Color("#A9B7C6"),

		InputBg: lipgloss.Color("#313335"),

		ListBg: lipgloss.Color("#1E1E1E"),
	}

	System = Theme{
		ID:   SystemID,
		Name: "System",

		Background: lipgloss.Color("#1e1e2e"),
		Foreground: lipgloss.Color("#cdd6f4"),

		Border:     lipgloss.Color("#45475a"),
		HelpFooter: lipgloss.Color("#6c7086"),
		LogoColor1: lipgloss.Color("#89b4fa"),
		LogoColor2: lipgloss.Color("#f5c2e7"),

		TaskTitle:    lipgloss.Color("#cdd6f4"),
		TaskPending:  lipgloss.Color("#a6adc8"),
		TaskDone:     lipgloss.Color("#585b70"),
		TaskDoneText: lipgloss.Color("#45475a"),

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
		ListBg:  lipgloss.Color("#1c1c28"),
	}

	Themes = []Theme{
		System,
		Catppuccin,
		TokyoNight,
		OneDark,
		Monochrome,
		Nord,
		Gruvbox,
		Ristretto,
		Monokai,
		Darcula,
	}
)

// ThemeID returns the config id for a theme.
func ThemeID(t Theme) string {
	if t.ID != "" {
		return t.ID
	}
	return Catppuccin.ID
}

// ByID returns theme for config id.
func ByID(id string) (Theme, int) {
	if IsSystem(id) {
		for i, th := range Themes {
			if th.ID == SystemID {
				return LoadSystem(), i
			}
		}
	}
	for i, th := range Themes {
		if th.ID == id {
			return th, i
		}
	}
	return Catppuccin, 1
}
