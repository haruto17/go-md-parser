package main

import (
	"fmt"

	"github.com/haruto17/go-md-parser/internal/parser"
)

func main() {
	input := "# h1\n\nThis is *Markdown* with **strong** and `code`."
	html := parser.Parse(input)
	fmt.Println(html)
}
