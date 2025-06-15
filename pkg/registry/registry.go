package registry

import (
  "fmt"
  "os"
  "path/filepath"
)

// Fetch will resolve tplSource (local path or git repo@ref) and return a filesystem path.
func Fetch(source, name string) (string, error) {
  // For v0: if source is a local folder, just join it:
  candidate := filepath.Join(source, name)
  if stat, err := os.Stat(candidate); err == nil && stat.IsDir() {
    return candidate, nil
  }
  return "", fmt.Errorf("template %q not found in %s", name, source)
}
