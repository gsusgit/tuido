package view

import (
	"testing"

	"github.com/gsus/todo-app/tui/internal/theme"
)

func TestPickLogoTier(t *testing.T) {
	if pickLogoTier(60) != logoTierFull {
		t.Fatalf("expected full at 60, got %v", pickLogoTier(60))
	}
	if pickLogoTier(30) != logoTierCompact {
		t.Fatalf("expected compact at 30, got %v", pickLogoTier(30))
	}
	if pickLogoTier(10) != logoTierMicro {
		t.Fatalf("expected micro at 10, got %v", pickLogoTier(10))
	}
}

func TestRenderLogoFitsWidth(t *testing.T) {
	th := theme.Themes[0]
	for _, w := range []int{80, 50, 35, 25} {
		logo := renderLogo(th, w, 24, w-statsColWidth-1)
		if logo == "" {
			t.Fatalf("empty logo at width %d", w)
		}
		if mw := maxLineWidth(logo); mw > w {
			t.Fatalf("logo width %d exceeds max %d at term %d", mw, w, w)
		}
	}
}

func TestMicroLogoNoHyphen(t *testing.T) {
	th := theme.Themes[0]
	logo := renderLogo(th, 15, 10, 12)
	if logo == "" {
		t.Fatal("expected micro logo")
	}
}

func TestLogoFullSplitLeavesDOIntact(t *testing.T) {
	wantDO := []string{"___  ____ ", " _ \\/ __ \\", "// / /_/ /", "____/\\____/ "}
	for i, line := range logoFullLines {
		split := logoFullSplit[i]
		got := line[split:]
		if got != wantDO[i] {
			t.Fatalf("line %d DO %q, want %q", i, got, wantDO[i])
		}
	}
}

func TestEffectiveMin(t *testing.T) {
	if effectiveMin(100, 5) >= 100 {
		t.Fatal("short height should reduce effective min")
	}
	if effectiveMin(20, 40) != 20 {
		t.Fatal("narrow width should be effective min")
	}
}
