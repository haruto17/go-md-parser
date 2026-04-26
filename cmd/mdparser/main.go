package main

import (
	"fmt"

	"github.com/haruto17/go-md-parser/internal/parser"
)

func main() {
	input := "# Hello\n\nThis is *Markdown*."
	html := parser.Parse(input)
	fmt.Println(html)
}
