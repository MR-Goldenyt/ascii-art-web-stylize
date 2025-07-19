package helpers

import (
	"strings"
)

// splitSegments chops a string like "Hello\nWorld"
// into ["Hello", "World"]
func SplitSegments(s string) []string {
    // Normalize both literal "\n" and actual newlines
    return strings.Split(s, "\n")
}
