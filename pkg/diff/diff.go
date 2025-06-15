package diff

import (
	"fmt"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

// Unified returns a git-style unified diff between oldText and newText,
// labeling the “from” file as a/<name> and the “to” file as b/<name>.
func Unified(name, oldText, newText string) string {
  // Compute the Myers edits between the two versions
  nameNew := span.URI(name)
  edits := myers.ComputeEdits(nameNew, oldText, newText)

  // Format them as a unified diff
  ud := gotextdiff.ToUnified("a/"+name, "b/"+name, oldText, edits)
  return fmt.Sprint(ud)
}
