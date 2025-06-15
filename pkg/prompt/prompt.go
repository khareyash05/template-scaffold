package prompt

import (
	"fmt"
)

// Collect reads a questions.yaml (or similar) from tplDir and interactively
// asks the user for each variable, returning a map of values.
func Collect(tplDir string) (map[string]interface{}, error) {
	// For now, a very simple prototype: look for questions.yaml,
	// but just hard-code one question as a demo.
	//
	// Later you’ll parse a YAML like:
	//   - name: ProjectName
	//     prompt: "Project name"
	//     default: "my-app"

	// Example stub:
	vars := make(map[string]interface{})
	fmt.Print("ProjectName (default: my-app): ")
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		// on empty or error, fall back to default
		input = "my-app"
	}
	vars["ProjectName"] = input
	return vars, nil
}
