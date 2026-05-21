package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSortTasksPendingFirst(t *testing.T) {
	tasks := []Task{
		{ID: "1", Title: "done", Completed: true, CreatedAt: time.Now()},
		{ID: "2", Title: "todo", Completed: false, CreatedAt: time.Now().Add(-time.Hour)},
	}
	SortTasks(tasks)
	if tasks[0].ID != "2" {
		t.Fatalf("pending task should be first, got %s", tasks[0].ID)
	}
}

func TestNewIDUnique(t *testing.T) {
	a, err := NewID()
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewID()
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatal("ids should differ")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	tasks := []Task{{ID: "abc", Title: "test", Category: CategoryPersonal, Priority: PriorityMedium, CreatedAt: time.Now()}}
	if err := Save(tasks); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 1 || loaded[0].Title != "test" {
		t.Fatalf("unexpected loaded: %+v", loaded)
	}
}

func TestLoadMissingFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	loaded, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 0 {
		t.Fatalf("expected empty slice, got %d tasks", len(loaded))
	}
}

func TestMigrateFromLegacy(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	legacyDir := filepath.Join(home, ".config", "todotui")
	if err := os.MkdirAll(legacyDir, 0o755); err != nil {
		t.Fatal(err)
	}
	legacyData := filepath.Join(legacyDir, "data.json")
	if err := os.WriteFile(legacyData, []byte(`{"tasks":[{"id":"legacy1","title":"Old task","completed":false,"priority":"media","category":"personal","created_at":"2024-01-01T00:00:00Z"}]}`), 0o644); err != nil {
		t.Fatal(err)
	}

	migrated, err := MigrateFromLegacy()
	if err != nil {
		t.Fatal(err)
	}
	if !migrated {
		t.Fatal("expected migration to occur")
	}

	newPath, err := dataPath()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(newPath); err != nil {
		t.Fatalf("expected migrated file at %s: %v", newPath, err)
	}
	if _, err := os.Stat(legacyData); !os.IsNotExist(err) {
		t.Fatal("legacy file should be renamed away")
	}

	tasks, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 1 || tasks[0].Title != "Old task" {
		t.Fatalf("unexpected tasks after migrate: %+v", tasks)
	}
}
