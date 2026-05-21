package view

import (
	"testing"

	"github.com/gsusgit/tuido/internal/config"
	"github.com/gsusgit/tuido/internal/model"
	"github.com/gsusgit/tuido/internal/storage"
)

func TestListContentRows(t *testing.T) {
	m := model.New(config.Config{})
	m.Tasks = []storage.Task{{ID: "1"}, {ID: "2"}, {ID: "3"}}
	if got := listContentRows(&m); got != 3 {
		t.Fatalf("expected 3 rows, got %d", got)
	}
	m.SearchActive = true
	if got := listContentRows(&m); got != 4 {
		t.Fatalf("expected 4 rows with search, got %d", got)
	}
}
