package view

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	_ "github.com/gsus/todo-app/tui/internal/i18n/locales"
	"github.com/gsus/todo-app/tui/internal/model"
	"github.com/gsus/todo-app/tui/internal/theme"
)

func TestTerminalSizeIsValid(t *testing.T) {
	if (TerminalSize{60, 20}).IsValid() != true {
		t.Fatal("60x20 should be valid at minimum")
	}
	if (TerminalSize{59, 20}).IsValid() != false {
		t.Fatal("59x20 should be invalid")
	}
	if (TerminalSize{60, 19}).IsValid() != false {
		t.Fatal("60x19 should be invalid")
	}
	if (TerminalSize{0, 0}).IsValid() != false {
		t.Fatal("0x0 should be invalid before first measure")
	}
}

func TestRenderTerminalTooSmallVisible(t *testing.T) {
	out := RenderTerminalTooSmall(theme.Themes[0], TerminalSize{Width: 50, Height: 20})
	if strings.TrimSpace(out) == "" {
		t.Fatal("fallback should not be empty")
	}
	if lipgloss.Height(out) < 20 {
		t.Fatalf("fallback height %d, want at least 20", lipgloss.Height(out))
	}
	if !strings.Contains(out, "60") {
		t.Fatalf("expected minimum width in output: %q", out)
	}
}

func TestRenderViewportGates(t *testing.T) {
	small := modelWithSize(40, 20)
	out := RenderViewport(&small, func(m *model.Model) string {
		return "MAIN"
	})
	if out == "MAIN" {
		t.Fatal("expected fallback, got main UI")
	}
	if strings.TrimSpace(out) == "" {
		t.Fatal("expected non-empty fallback output")
	}
	if !strings.Contains(out, "60") || !strings.Contains(out, "20") {
		t.Fatalf("fallback should mention minimum 60x20, got: %q", out)
	}

	large := modelWithSize(60, 20)
	if got := RenderViewport(&large, func(m *model.Model) string { return "MAIN" }); got != "MAIN" {
		t.Fatalf("expected main UI, got %q", got)
	}
}

func modelWithSize(w, h int) model.Model {
	return model.Model{
		Width:  w,
		Height: h,
		Theme:  theme.Themes[0],
	}
}
