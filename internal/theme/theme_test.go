package theme

import "testing"

func TestThemesCount(t *testing.T) {
	if len(Themes) < 10 {
		t.Fatalf("expected at least 10 themes, got %d", len(Themes))
	}
}

func TestByIDKnown(t *testing.T) {
	cases := []struct {
		id   string
		want string
	}{
		{"system", "system"},
		{"catppuccin", "catppuccin"},
		{"tokyo-night", "tokyo-night"},
		{"nord", "nord"},
		{"gruvbox", "gruvbox"},
		{"ristretto", "ristretto"},
		{"monokai", "monokai"},
		{"darcula", "darcula"},
	}
	for _, tc := range cases {
		th, idx := ByID(tc.id)
		if th.ID != tc.want {
			t.Fatalf("ByID(%q).ID = %q, want %q", tc.id, th.ID, tc.want)
		}
		if Themes[idx].ID != tc.want {
			t.Fatalf("Themes[%d].ID = %q", idx, Themes[idx].ID)
		}
	}
}

func TestByIDUnknown(t *testing.T) {
	th, idx := ByID("unknown-theme")
	if th.ID != Catppuccin.ID {
		t.Fatalf("unknown id: got theme %q, want %q", th.ID, Catppuccin.ID)
	}
	if idx != 1 {
		t.Fatalf("unknown id: got index %d, want 1", idx)
	}
}

func TestThemeIDMatchesSlice(t *testing.T) {
	for _, th := range Themes {
		if ThemeID(th) != th.ID {
			t.Fatalf("ThemeID(%q) = %q, want %q", th.Name, ThemeID(th), th.ID)
		}
		if th.ID == "" {
			t.Fatalf("theme %q has empty ID", th.Name)
		}
	}
}
