package theme

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSystemFromOmarchy(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("no home dir")
	}
	colorsPath := filepath.Join(home, omarchyColorsRel)
	if _, err := os.Stat(colorsPath); err != nil {
		t.Skip("omarchy colors.toml not present")
	}

	th := LoadSystem()
	if th.ID != SystemID {
		t.Fatalf("LoadSystem().ID = %q, want %q", th.ID, SystemID)
	}
	if th.Name != "System" {
		t.Fatalf("LoadSystem().Name = %q, want System", th.Name)
	}
	if string(th.Background) == "" || string(th.Foreground) == "" {
		t.Fatal("expected non-empty background and foreground")
	}
}

func TestByIDSystem(t *testing.T) {
	th, idx := ByID(SystemID)
	if th.ID != SystemID {
		t.Fatalf("ByID(system).ID = %q", th.ID)
	}
	if Themes[idx].ID != SystemID {
		t.Fatalf("Themes[%d].ID = %q, want system", idx, Themes[idx].ID)
	}
}

func TestOmarchySignature(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("no home dir")
	}
	if _, err := os.Stat(filepath.Join(home, omarchyColorsRel)); err != nil {
		t.Skip("omarchy colors.toml not present")
	}
	sig, ok := OmarchySignature()
	if !ok || sig == "" {
		t.Fatal("expected valid omarchy signature")
	}
}

func TestParseColorsTOML(t *testing.T) {
	data := []byte(`background = "#111422"
foreground = "#bcc1dc"
accent = "#69c3ff"
color1 = "#e35535"
`)
	got, err := parseColorsTOML(data)
	if err != nil {
		t.Fatal(err)
	}
	if got["background"] != "#111422" {
		t.Fatalf("background = %q", got["background"])
	}
	if got["color1"] != "#e35535" {
		t.Fatalf("color1 = %q", got["color1"])
	}
}
