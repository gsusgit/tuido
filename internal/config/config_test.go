package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefault(t *testing.T) {
	cfg := Default()
	if cfg.Lang != DefaultLang {
		t.Fatalf("lang: got %q want %q", cfg.Lang, DefaultLang)
	}
	if cfg.Theme != DefaultTheme {
		t.Fatalf("theme: got %q want %q", cfg.Theme, DefaultTheme)
	}
}

func TestLoadMissingFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Lang != DefaultLang || cfg.Theme != DefaultTheme {
		t.Fatalf("got %+v, want defaults", cfg)
	}
}

func TestLoadEmptyFields(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	dir := filepath.Join(home, ".config", "tuido")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(`{"lang":"","theme":""}`), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Lang != DefaultLang || cfg.Theme != DefaultTheme {
		t.Fatalf("got %+v, want defaults filled", cfg)
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	dir := filepath.Join(home, ".config", "tuido")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(`not json`), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid config.json")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	want := Config{Lang: "es", Theme: "nord"}
	if err := Save(want); err != nil {
		t.Fatal(err)
	}
	got, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if got.Lang != "es" || got.Theme != "nord" {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
