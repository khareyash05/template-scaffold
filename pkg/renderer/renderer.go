package renderer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/khareyash05/scaffold-templates/pkg/diff"
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
		data, err := os.ReadFile(path)
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

		// strip .tmpl extension if present
		if strings.HasSuffix(outPath, ".tmpl") {
			outPath = strings.TrimSuffix(outPath, ".tmpl")
		}

		// write
		return os.WriteFile(outPath, buf.Bytes(), info.Mode())
	})
}

func DryRun(tplDir string, vars map[string]interface{}, outDir string, diffMode bool) error {
	return filepath.Walk(tplDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		rel, _ := filepath.Rel(tplDir, path)

		// Render the template into newBuf
		data, _ := os.ReadFile(path)
		tmpl := template.Must(template.New(rel).Parse(string(data)))
		var newBuf bytes.Buffer
		_ = tmpl.Execute(&newBuf, vars)
		newText := newBuf.String()

		// strip .tmpl extension if present
		outRel := rel
		if strings.HasSuffix(outRel, ".tmpl") {
			outRel = strings.TrimSuffix(outRel, ".tmpl")
		}

		if diffMode {
			// Try to read existing file (could be empty/nonexistent)
			existingPath := filepath.Join(outDir, outRel)
			oldBytes, _ := os.ReadFile(existingPath)
			oldText := string(oldBytes)

			fmt.Println(diff.Unified(outRel, oldText, newText))
		} else {
			fmt.Printf("[DRY-RUN] would write %s\n", outRel)
		}

		return nil
	})
}
