package discover

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// GetTodoFilePath returns the path to .gotodo.json based on Git repo root or current dir.
// If inside a Git repo, returns .gotodo.json at the repo root.
// Otherwise, returns .gotodo.json in the current directory.
func GetTodoFilePath(cwd string) string {
	// Try to get Git root
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = cwd
	output, err := cmd.Output()
	if err == nil {
		gitRoot := strings.TrimSpace(string(output))
		return filepath.Join(gitRoot, ".gotodo.json")
	}
	// Not in Git repo, use current dir
	return filepath.Join(cwd, ".gotodo.json")
}
