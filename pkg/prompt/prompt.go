package prompt

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strconv"
  "strings"

  "gopkg.in/yaml.v3"
)

// Question represents one prompt in questions.yaml
type Question struct {
  Name    string      `yaml:"name"`
  Prompt  string      `yaml:"prompt"`
  Default interface{} `yaml:"default"`
  // Optional override; if empty we’ll infer from Default
  Type    string      `yaml:"type,omitempty"` 
}

// Collect parses questions.yaml from tplDir and interactively asks each one,
// inferring or using explicit types for answers.
func Collect(tplDir string) (map[string]interface{}, error) {
  qFile := filepath.Join(tplDir, "questions.yaml")
  data, err := ioutil.ReadFile(qFile)
  if err != nil {
    return nil, fmt.Errorf("reading %s: %w", qFile, err)
  }

  var qs []Question
  if err := yaml.Unmarshal(data, &qs); err != nil {
    return nil, fmt.Errorf("parsing %s: %w", qFile, err)
  }

  vars := make(map[string]interface{}, len(qs))
  reader := bufio.NewReader(os.Stdin)

  for _, q := range qs {
    defVal := q.Default
    defStr := fmt.Sprintf("%v", defVal)

    // Determine the type: explicit override or infer from Default
    typ := q.Type
    if typ == "" {
      switch defVal.(type) {
      case bool:
        typ = "bool"
      case int, int64, float64:
        typ = "number"
      default:
        typ = "string"
      }
    }

    // Prompt
    fmt.Printf("%s (default: %s): ", q.Prompt, defStr)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    // Use default if empty
    if input == "" {
      vars[q.Name] = defVal
      continue
    }

    // Parse according to type
    switch typ {
    case "bool":
      // accept y/n or true/false
      low := strings.ToLower(input)
      if low == "y" || low == "yes" {
        vars[q.Name] = true
      } else if low == "n" || low == "no" {
        vars[q.Name] = false
      } else {
        // try strconv.ParseBool
        b, err := strconv.ParseBool(low)
        if err != nil {
          return nil, fmt.Errorf("invalid bool for %s: %q", q.Name, input)
        }
        vars[q.Name] = b
      }

    case "number":
      // try integer first
      if i, err := strconv.Atoi(input); err == nil {
        vars[q.Name] = i
      } else if f, err := strconv.ParseFloat(input, 64); err == nil {
        vars[q.Name] = f
      } else {
        return nil, fmt.Errorf("invalid number for %s: %q", q.Name, input)
      }

    default: // string
      vars[q.Name] = input
    }
  }

  return vars, nil
}
