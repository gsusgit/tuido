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
	os.Setenv("HOME", home)
	// storage uses UserHomeDir
	old := os.Getenv("HOME")
	defer os.Setenv("HOME", old)

	dir := filepath.Join(home, ".config", "tuido")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}

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
