package main

import (
	"fmt"

	"github.com/haruto17/go-md-parser/internal/parser"
)

func main() {
	input := "# h1\n\n## h2\n\n### h3\n\n#### h4\n\n##### h5\n\n###### h6\n\nThis is *Markdown*."
	html := parser.Parse(input)
	fmt.Println(html)
}
