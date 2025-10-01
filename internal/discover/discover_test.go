package discover

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetTodoFilePath(t *testing.T) {
	// Test in a Git repo
	dir, err := ioutil.TempDir("", "gotodo-git-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Initialize a Git repo
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command("git", "init").Run(); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command("git", "config", "user.email", "test@example.com").Run(); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command("git", "config", "user.name", "Test User").Run(); err != nil {
		t.Fatal(err)
	}

	// Create a subdirectory
	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatal(err)
	}

	// Get the actual git root
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = subdir
	gitOutput, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	gitRoot := strings.TrimSpace(string(gitOutput))

	path := GetTodoFilePath(subdir)
	expected := filepath.Join(gitRoot, ".gotodo.json")
	if path != expected {
		t.Fatalf("expected %s, got %s", expected, path)
	}

	// Test not in Git repo
	nonGitDir, err := ioutil.TempDir("", "gotodo-nongit-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(nonGitDir)

	path2 := GetTodoFilePath(nonGitDir)
	expected2 := filepath.Join(nonGitDir, ".gotodo.json")
	if path2 != expected2 {
		t.Fatalf("expected %s, got %s", expected2, path2)
	}
}
