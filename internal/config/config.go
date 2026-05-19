package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// Config holds persistent user preferences.
type Config struct {
	Lang  string `json:"lang"`
	Theme string `json:"theme"`
}

// Default returns config with system language and default theme.
func Default() Config {
	return Config{
		Lang:  detectSystemLang(),
		Theme: "catppuccin",
	}
}

// Load reads ~/.config/tuido/config.json.
func Load() (Config, error) {
	path, err := configPath()
	if err != nil {
		return Default(), err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return Default(), err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Default(), err
	}
	if cfg.Lang == "" {
		cfg.Lang = Default().Lang
	}
	if cfg.Theme == "" {
		cfg.Theme = Default().Theme
	}
	return cfg, nil
}

// Save writes config atomically.
func Save(cfg Config) error {
	dir, err := configDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "tuido"), nil
}

func configPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func detectSystemLang() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		return "en"
	}
	lang = strings.ToLower(lang)
	switch {
	case strings.HasPrefix(lang, "es"):
		return "es"
	case strings.HasPrefix(lang, "fr"):
		return "fr"
	case strings.HasPrefix(lang, "de"):
		return "de"
	case strings.HasPrefix(lang, "it"):
		return "it"
	case strings.HasPrefix(lang, "pt"):
		return "pt"
	default:
		return "en"
	}
}
