package parser

import (
	"strings"
)

func Parse(input string) string {
	lines := strings.Split(input, "\n")
	var out []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if content, ok := strings.CutPrefix(line, "# "); ok {
			out = append(out, "<h1>"+content+"</h1>")
			continue
		}

		out = append(out, "<p>"+line+"</p>")
	}

	return strings.Join(out, "\n")
}
