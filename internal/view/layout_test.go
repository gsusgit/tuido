package view

import (
	"strings"
	"testing"
)

func TestClipToMaxHeight(t *testing.T) {
	s := "a\nb\nc\nd\ne"
	clipped := clipToMaxHeight(s, 3)
	if lineCount(clipped) > 3 {
		t.Fatalf("expected at most 3 lines, got %d", lineCount(clipped))
	}
}

func TestRenderShellHeight(t *testing.T) {
	header := "head\nline2"
	body := "body1\nbody2\nbody3\nbody4\nbody5"
	footer := "foot"
	out := renderShell(40, 8, header, body, footer)
	// Footer must be present in output
	if !strings.Contains(out, "foot") {
		t.Fatal("footer should appear in shell output")
	}
	if lineCount(out) < 8 {
		t.Fatalf("shell should fill at least height 8, got %d lines", lineCount(out))
	}
}
