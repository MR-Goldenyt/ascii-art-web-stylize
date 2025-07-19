package helpers

import "strings"

func RenderAscii(segments []string, fontLines []string) (string, int) {
	// 6) Render ASCII art
	var out strings.Builder
	rowsPerChar := 9
	linecount := 0
	for i := 0; i < len(segments); {
		seg := segments[i]
		if seg != "" {
			// render 8 rows per non-blank segment
			for row := 0; row < 8; row++ {
				for _, rnr := range seg {
					idx := int(rnr) - 32
					start := idx*rowsPerChar + 1
					out.WriteString(fontLines[start+row])
				}
				out.WriteByte('\n')

			}
			linecount += 8
			i++
		} else {
			// blank line(s)
			j := i
			for j < len(segments) && segments[j] == "" {
				j++
			}
			for k := 0; k < j-i; k++ {
				out.WriteByte('\n')
				linecount++
			}
			i = j
		}
	}
	return out.String(), linecount+2
}
