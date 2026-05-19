package storage

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

)

type Priority string

const (
	PriorityHigh   Priority = "alta"
	PriorityMedium Priority = "media"
	PriorityLow    Priority = "baja"
)

type Category string

const (
	CategoryPersonal Category = "personal"
	CategoryTrabajo  Category = "trabajo"
)

// AllCategories returns selectable categories.
func AllCategories() []Category {
	return []Category{CategoryPersonal, CategoryTrabajo}
}

// AllPriorities returns selectable priorities.
func AllPriorities() []Priority {
	return []Priority{PriorityHigh, PriorityMedium, PriorityLow}
}

// Label returns a display label for the category (i18n keys handled in view).
func (c Category) String() string { return string(c) }

// Label returns priority code.
func (p Priority) String() string { return string(p) }

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	Priority    Priority   `json:"priority"`
	Category    Category   `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type dataFile struct {
	Tasks []Task `json:"tasks"`
}

func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".config", "tuido")
	return dir, os.MkdirAll(dir, 0o755)
}

func dataPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "data.json"), nil
}

// MigrateFromLegacy moves ~/.config/todotui/data.json to tuido if needed.
func MigrateFromLegacy() (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}
	oldPath := filepath.Join(home, ".config", "todotui", "data.json")
	newPath, err := dataPath()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(newPath); err == nil {
		return false, nil
	}
	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return false, nil
	}
	if err := os.MkdirAll(filepath.Dir(newPath), 0o755); err != nil {
		return false, err
	}
	if err := os.Rename(oldPath, newPath); err != nil {
		return false, fmt.Errorf("migrate data: %w", err)
	}
	return true, nil
}

// DataFilePath returns the path to data.json for CLI messages.
func DataFilePath() (string, error) {
	return dataPath()
}

// Load reads tasks from ~/.config/tuido/data.json.
func Load() ([]Task, error) {
	path, err := dataPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var df dataFile
	if err := json.Unmarshal(data, &df); err != nil {
		return nil, fmt.Errorf("parse data.json: %w", err)
	}
	SortTasks(df.Tasks)
	return df.Tasks, nil
}

// Save writes tasks to ~/.config/tuido/data.json.
func Save(tasks []Task) error {
	path, err := dataPath()
	if err != nil {
		return err
	}
	sorted := make([]Task, len(tasks))
	copy(sorted, tasks)
	SortTasks(sorted)
	df := dataFile{Tasks: sorted}
	data, err := json.MarshalIndent(df, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

// Reset removes the data file.
func Reset() error {
	path, err := dataPath()
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// SortTasks sorts pending first, then by created date descending.
func SortTasks(tasks []Task) {
	sort.SliceStable(tasks, func(i, j int) bool {
		if tasks[i].Completed != tasks[j].Completed {
			return !tasks[i].Completed
		}
		return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
	})
}

// NewID generates a random unique task ID.
func NewID() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
