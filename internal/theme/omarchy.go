package theme

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const SystemID = "system"

// Omarchy paths (Hyprland/Omarchy desktop theme).
var (
	omarchyColorsRel    = filepath.Join(".config", "omarchy", "current", "theme", "colors.toml")
	omarchyThemeNameRel = filepath.Join(".config", "omarchy", "current", "theme.name")
)

type omarchyColors struct {
	Background          string
	Foreground          string
	Accent              string
	SelectionBackground string
	SelectionForeground string
	Color0              string
	Color1              string
	Color2              string
	Color3              string
	Color5              string
	Color8              string
	Color15             string
}

// IsSystem reports whether id is the dynamic Omarchy/Hyprland theme.
func IsSystem(id string) bool {
	return id == SystemID
}

// LoadSystem reads the active Omarchy theme and maps it to a TUI palette.
func LoadSystem() Theme {
	colors, ok := loadOmarchyColors()
	if !ok {
		fb := Catppuccin
		fb.ID = SystemID
		fb.Name = "System"
		return fb
	}
	return mapOmarchyColors(colors)
}

// OmarchySignature returns a hash that changes when the Omarchy theme or its colors change.
func OmarchySignature() (string, bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", false
	}

	colorsPath := filepath.Join(home, omarchyColorsRel)
	namePath := filepath.Join(home, omarchyThemeNameRel)

	colorsInfo, err := os.Stat(colorsPath)
	if err != nil {
		return "", false
	}

	themeName, _ := os.ReadFile(namePath)
	h := sha256.New()
	h.Write(themeName)
	h.Write([]byte{0})
	h.Write([]byte(colorsInfo.ModTime().UTC().Format("2006-01-02T15:04:05")))
	h.Write([]byte{0})
	if data, err := os.ReadFile(colorsPath); err == nil {
		h.Write(data)
	}
	return hex.EncodeToString(h.Sum(nil)), true
}

func loadOmarchyColors() (omarchyColors, bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		return omarchyColors{}, false
	}
	path := filepath.Join(home, omarchyColorsRel)
	data, err := os.ReadFile(path)
	if err != nil {
		return omarchyColors{}, false
	}
	parsed, err := parseColorsTOML(data)
	if err != nil {
		return omarchyColors{}, false
	}
	c := omarchyColors{
		Background:          parsed["background"],
		Foreground:          parsed["foreground"],
		Accent:              parsed["accent"],
		SelectionBackground: parsed["selection_background"],
		SelectionForeground: parsed["selection_foreground"],
		Color0:              parsed["color0"],
		Color1:              parsed["color1"],
		Color2:              parsed["color2"],
		Color3:              parsed["color3"],
		Color5:              parsed["color5"],
		Color8:              parsed["color8"],
		Color15:             parsed["color15"],
	}
	if c.Background == "" || c.Foreground == "" {
		return omarchyColors{}, false
	}
	if c.Accent == "" {
		c.Accent = firstNonEmpty(parsed["color4"], c.Foreground)
	}
	if c.Color0 == "" {
		c.Color0 = c.Background
	}
	if c.Color8 == "" {
		c.Color8 = adjustHex(c.Background, -0.12)
	}
	if c.Color15 == "" {
		c.Color15 = c.Foreground
	}
	if c.Color1 == "" {
		c.Color1 = "#f38ba8"
	}
	if c.Color2 == "" {
		c.Color2 = "#a6e3a1"
	}
	if c.Color3 == "" {
		c.Color3 = "#fab387"
	}
	if c.Color5 == "" {
		c.Color5 = firstNonEmpty(parsed["color13"], c.Accent)
	}
	if c.SelectionBackground == "" {
		c.SelectionBackground = c.Color0
	}
	if c.SelectionForeground == "" {
		c.SelectionForeground = c.Foreground
	}
	return c, true
}

func mapOmarchyColors(c omarchyColors) Theme {
	bg := lipgloss.Color(c.Background)
	fg := lipgloss.Color(c.Foreground)
	muted := lipgloss.Color(c.Color8)

	return Theme{
		ID:   SystemID,
		Name: "System",

		Background: bg,
		Foreground: fg,

		Border:     muted,
		HelpFooter: muted,
		LogoColor1: lipgloss.Color(c.Accent),
		LogoColor2: lipgloss.Color(c.Color5),

		TaskTitle:    fg,
		TaskPending:  lipgloss.Color(c.Color15),
		TaskDone:     muted,
		TaskDoneText: lipgloss.Color(adjustHex(c.Color8, -0.18)),

		PriorityHigh:   lipgloss.Color(c.Color1),
		PriorityMedium: lipgloss.Color(c.Color3),
		PriorityLow:    lipgloss.Color(c.Color2),

		Accent:  lipgloss.Color(c.Accent),
		Accent2: lipgloss.Color(c.Color5),

		SelectedBg: lipgloss.Color(c.SelectionBackground),
		SelectedFg: lipgloss.Color(c.SelectionForeground),

		ModalBg: lipgloss.Color(adjustHex(c.Background, 0.06)),
		ModalFg: fg,
		Overlay: lipgloss.Color("rgba(0,0,0,0.6)"),

		HotkeyBg: lipgloss.Color(c.Color0),
		HotkeyFg: fg,

		InputBg: lipgloss.Color(adjustHex(c.Background, -0.04)),
		ListBg:  lipgloss.Color(adjustHex(c.Background, -0.10)),
	}
}

func parseColorsTOML(data []byte) (map[string]string, error) {
	out := make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		val = strings.Trim(val, `"`)
		out[key] = val
	}
	return out, nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func adjustHex(hex string, amount float64) string {
	c, err := colorful.Hex(hex)
	if err != nil {
		return hex
	}
	h, s, l := c.Hsl()
	l = clamp01(l + amount)
	return colorful.Hsl(h, s, l).Hex()
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
