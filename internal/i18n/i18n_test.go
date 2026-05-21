package i18n_test

import (
	"testing"

	"github.com/gsusgit/tuido/internal/i18n"
	_ "github.com/gsusgit/tuido/internal/i18n/locales"
)

// allUIKeys mirrors exported Key constants in i18n.go for catalog parity checks.
var allUIKeys = []i18n.Key{
	i18n.KeyNewTask,
	i18n.KeyFilter,
	i18n.KeyFilterActive,
	i18n.KeyEdit,
	i18n.KeyDelete,
	i18n.KeyComplete,
	i18n.KeyNavigate,
	i18n.KeyTheme,
	i18n.KeyThemeTitle,
	i18n.KeyExit,
	i18n.KeySearch,
	i18n.KeyEmptyList,
	i18n.KeyEmptyFiltered,
	i18n.KeyTitleRequired,
	i18n.KeyCompletedOf,
	i18n.KeyScrollMore,
	i18n.KeyTitle,
	i18n.KeyCategory,
	i18n.KeyPriority,
	i18n.KeyStatus,
	i18n.KeyPending,
	i18n.KeyStatusCompleted,
	i18n.KeyFilterTitle,
	i18n.KeyAll,
	i18n.KeyTabNext,
	i18n.KeyEnterSave,
	i18n.KeyEnterApply,
	i18n.KeyEscCancel,
	i18n.KeyBack,
	i18n.KeyReset,
	i18n.KeyTerminalSmall,
	i18n.KeyWindowTooSmall,
	i18n.KeyMinimumSizeRequired,
	i18n.KeyDeleteConfirm,
	i18n.KeyDeletePending,
	i18n.KeyTaskCreated,
	i18n.KeyTaskSaved,
	i18n.KeyTaskDeleted,
	i18n.KeySaveError,
	i18n.KeyLangChanged,
	i18n.KeyResetPrompt,
	i18n.KeyResetDone,
	i18n.KeyResetCancelled,
	i18n.KeyResetInvalid,
	i18n.KeyHelpUsage,
	i18n.KeyMigrating,
	i18n.KeyPersonal,
	i18n.KeyTrabajo,
	i18n.KeyPriorityHigh,
	i18n.KeyPriorityMedium,
	i18n.KeyPriorityLow,
	i18n.KeySelectLang,
	i18n.KeyLangNav,
	i18n.KeyControls,
	i18n.KeyHelpMenu,
	i18n.KeySort,
	i18n.KeySortField,
	i18n.KeySortDirection,
	i18n.KeySortDefault,
	i18n.KeySortAsc,
	i18n.KeySortDesc,
}

func TestEnglishCatalogNonEmpty(t *testing.T) {
	i18n.SetLang("en")
	for _, key := range allUIKeys {
		msg := i18n.T(key)
		if msg == "" || msg == string(key) {
			t.Fatalf("en missing translation for %q: %q", key, msg)
		}
	}
}

func TestCatalogParity(t *testing.T) {
	langs := []string{"es", "fr", "de", "it", "pt"}
	for _, lang := range langs {
		if !i18n.IsSupported(lang) {
			t.Fatalf("expected supported language %q", lang)
		}
		i18n.SetLang(lang)
		for _, key := range allUIKeys {
			msg := i18n.T(key)
			if msg == "" || msg == string(key) {
				t.Fatalf("%s missing translation for %q", lang, key)
			}
		}
	}
	i18n.SetLang("en")
}

func TestSetLangUnsupported(t *testing.T) {
	i18n.SetLang("en")
	before := i18n.Current()
	i18n.SetLang("xx")
	if i18n.Current() != before {
		t.Fatalf("unsupported lang should not change current: got %q", i18n.Current())
	}
}

func TestSetLangSupported(t *testing.T) {
	i18n.SetLang("es")
	if i18n.Current() != "es" {
		t.Fatalf("got %q, want es", i18n.Current())
	}
	i18n.SetLang("en")
}

func TestThemeTitleNotRawKey(t *testing.T) {
	i18n.SetLang("en")
	msg := i18n.T(i18n.KeyThemeTitle)
	if msg == string(i18n.KeyThemeTitle) {
		t.Fatalf("expected translated theme title, got raw key %q", msg)
	}
}
