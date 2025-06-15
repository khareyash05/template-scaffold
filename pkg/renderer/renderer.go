package renderer

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "text/template"
)

// Render walks tplDir, applies vars to every file, writes into outDir.
func Render(tplDir string, vars map[string]interface{}, outDir string) error {
  return filepath.Walk(tplDir, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    rel, _ := filepath.Rel(tplDir, path)
    outPath := filepath.Join(outDir, rel)

    if info.IsDir() {
      return os.MkdirAll(outPath, 0755)
    }

    // read & parse
    data, err := ioutil.ReadFile(path)
    if err != nil {
      return err
    }
    tmpl, err := template.New(rel).Parse(string(data))
    if err != nil {
      return fmt.Errorf("parsing %s: %w", rel, err)
    }

    // execute into buffer
    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, vars); err != nil {
      return fmt.Errorf("rendering %s: %w", rel, err)
    }

    // write
    return ioutil.WriteFile(outPath, buf.Bytes(), info.Mode())
  })
}

// DryRun can show file-by-file diffs if diffMode=true, otherwise just list.
func DryRun(tplDir string, vars map[string]interface{}, outDir string, diffMode bool) error {
  // TODO: implement using pkg/diff
  fmt.Printf("Dry run: would scaffold %s into %s (diff: %v)\n", tplDir, outDir, diffMode)
  return nil
}
