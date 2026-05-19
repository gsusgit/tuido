package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/gsus/todo-app/tui/internal/theme"
)

type logoTier int

const (
	logoTierMicro logoTier = iota
	logoTierCompact
	logoTierFull
	statsColWidth = 9
)

// Official 4-line Slant ASCII (TUIDO).
var logoFullLines = []string{
	" ________  ______    ___  ____ ",
	"/_  __/ / / /  _/___/ _ \\/ __ \\",
	" / / / /_/ // //___/ // / /_/ /",
	"/_/  \\____/___/   /____/\\____/ ",
}

// Split index per line: TUI part | DO part (slant art boundary for D/O).
var logoFullSplit = []int{21, 21, 21, 19}

// Compact 2-line Slant — readable mini TUIDO.
var logoCompactLines = []struct{ tui, do string }{
	{" ______   ___ ", " ____ "},
	{"/_  _/___/ _ \\", "/ __ \\"},
}

const (
	logoTierFullMin    = 52 // only on wide/tall terminals
	logoTierCompactMin = 20
)

type logoLine struct {
	text  string
	split int // byte index where DO starts
}

// effectiveMin mimics CSS vmin: scale with the smaller of terminal width
// or height-derived width.
func effectiveMin(width, height int) int {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}
	hScale := height * 5
	if hScale < width {
		return hScale
	}
	return width
}

func pickLogoTier(vmin int) logoTier {
	switch {
	case vmin >= logoTierFullMin:
		return logoTierFull
	case vmin >= logoTierCompactMin:
		return logoTierCompact
	default:
		return logoTierMicro
	}
}

func maxLineWidth(s string) int {
	max := 0
	for _, line := range strings.Split(s, "\n") {
		if w := lipgloss.Width(line); w > max {
			max = w
		}
	}
	return max
}

func logoLinesForTier(tier logoTier) []logoLine {
	switch tier {
	case logoTierFull:
		out := make([]logoLine, len(logoFullLines))
		for i, line := range logoFullLines {
			split := logoFullSplit[i]
			if split > len(line) {
				split = len(line)
			}
			out[i] = logoLine{text: line, split: split}
		}
		return out
	case logoTierCompact:
		out := make([]logoLine, len(logoCompactLines))
		for i, l := range logoCompactLines {
			out[i] = logoLine{text: l.tui + l.do, split: len(l.tui)}
		}
		return out
	default:
		return []logoLine{{text: "TUIDO", split: 3}}
	}
}

func fitLogoLines(lines []logoLine, maxW int) []logoLine {
	if maxW < 1 {
		maxW = 1
	}
	out := make([]logoLine, 0, len(lines))
	for _, l := range lines {
		line := l.text
		if len(line) > maxW {
			line = truncatePlain(line, maxW)
		}
		if line == "" {
			continue
		}
		split := l.split
		if split > len(line) {
			split = len(line)
		}
		out = append(out, logoLine{text: line, split: split})
	}
	return out
}

func truncatePlain(s string, maxW int) string {
	if maxW < 1 || len(s) <= maxW {
		return s
	}
	return s[:maxW]
}

func colorizeLogoLines(lines []logoLine, t theme.Theme) string {
	tuiStyle := lipgloss.NewStyle().Foreground(t.LogoColor1).Bold(true)
	doStyle := lipgloss.NewStyle().Foreground(t.LogoColor2).Bold(true)

	styled := make([]string, 0, len(lines))
	for _, l := range lines {
		if l.split <= 0 {
			styled = append(styled, doStyle.Render(l.text))
			continue
		}
		if l.split >= len(l.text) {
			styled = append(styled, tuiStyle.Render(l.text))
			continue
		}
		styled = append(styled,
			tuiStyle.Render(l.text[:l.split])+doStyle.Render(l.text[l.split:]),
		)
	}
	return strings.Join(styled, "\n")
}

func plainLogoWidth(lines []logoLine) int {
	max := 0
	for _, l := range lines {
		if w := len(l.text); w > max {
			max = w
		}
	}
	return max
}

// renderLogo picks a Slant tier from terminal vmin, then downgrades until it fits maxWidth.
func renderLogo(t theme.Theme, width, height, maxWidth int) string {
	if maxWidth < 8 {
		maxWidth = 8
	}
	vmin := effectiveMin(width, height)
	tier := pickLogoTier(vmin)

	for tier >= logoTierMicro {
		lines := fitLogoLines(logoLinesForTier(tier), maxWidth)
		if len(lines) == 0 {
			tier--
			continue
		}
		if plainLogoWidth(lines) <= maxWidth {
			return colorizeLogoLines(lines, t)
		}
		tier--
	}

	lines := fitLogoLines(logoLinesForTier(logoTierMicro), maxWidth)
	if len(lines) == 0 {
		return colorizeLogoLines(logoLinesForTier(logoTierMicro), t)
	}
	return colorizeLogoLines(lines, t)
}
