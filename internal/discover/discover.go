package discover

import (
	"os"
	"path/filepath"
)

// FindNearestTodoFile walks up from startDir to the filesystem root looking for
// a ".gotodo.json" file. If found, it returns the absolute path. If not
// found, it returns an empty string and nil error.
func FindNearestTodoFile(startDir string) (string, error) {
	dir := startDir
	for {
		candidate := filepath.Join(dir, ".gotodo.json")
		if _, err := os.Stat(candidate); err == nil {
			abs, err := filepath.Abs(candidate)
			if err != nil {
				return candidate, nil
			}
			return abs, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// reached root
			return "", nil
		}
		dir = parent
	}
}
