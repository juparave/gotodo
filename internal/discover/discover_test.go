package discover

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFindNearestTodoFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "gotodo-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create nested dirs
	nested := filepath.Join(dir, "a", "b", "c")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}

	// place .gotodo.json in 'a'
	marker := filepath.Join(dir, "a", ".gotodo.json")
	if err := ioutil.WriteFile(marker, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}

	found, err := FindNearestTodoFile(nested)
	if err != nil {
		t.Fatal(err)
	}
	if found == "" {
		t.Fatalf("expected to find .gotodo.json, got none")
	}
	if filepath.Base(found) != ".gotodo.json" {
		t.Fatalf("unexpected file found: %s", found)
	}
}
