package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/juparave/gotodo/internal/model"
)

func TestOpenIndexes(t *testing.T) {
	dir, err := ioutil.TempDir("", "gotodo-store-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, ".gotodo.json")
	s := NewJSONFileStore(path)
	// initialize
	if err := s.Init(); err != nil {
		t.Fatal(err)
	}

	// add 3 todos: 2 open, 1 done
	s.Add(model.NewTodo("one"))
	s.Add(model.NewTodo("two"))
	t3 := model.NewTodo("three")
	t3.Done = true
	s.Add(t3)

	idxes := s.OpenIndexes()
	if len(idxes) != 2 {
		t.Fatalf("expected 2 open indexes, got %d", len(idxes))
	}
	// check that they point to the first two entries (0 and 1)
	if idxes[0] != 0 || idxes[1] != 1 {
		t.Fatalf("unexpected open indexes: %#v", idxes)
	}
}
