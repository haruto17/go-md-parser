package parser

import (
	"strings"
)

func Parse(input string) string {
	lines := strings.Split(input, "\n")
	var out []string

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		if content, ok := parseHeading(line); ok {
			out = append(out, content)
			continue
		}

		out = append(out, renderParagraph(line))
	}

	return strings.Join(out, "\n")
}

func parseHeading(line string) (string, bool) {
	switch {
	case strings.HasPrefix(line, "###### "):
		return "<h6>" + strings.TrimPrefix(line, "###### ") + "</h6>", true
	case strings.HasPrefix(line, "##### "):
		return "<h5>" + strings.TrimPrefix(line, "##### ") + "</h5>", true
	case strings.HasPrefix(line, "#### "):
		return "<h4>" + strings.TrimPrefix(line, "#### ") + "</h4>", true
	case strings.HasPrefix(line, "### "):
		return "<h3>" + strings.TrimPrefix(line, "### ") + "</h3>", true
	case strings.HasPrefix(line, "## "):
		return "<h2>" + strings.TrimPrefix(line, "## ") + "</h2>", true
	case strings.HasPrefix(line, "# "):
		return "<h1>" + strings.TrimPrefix(line, "# ") + "</h1>", true
	default:
		return "", false
	}
}
