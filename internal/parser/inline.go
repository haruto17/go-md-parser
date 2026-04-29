package parser

import "strings"

type InlineTokenType int

const (
	TokenText InlineTokenType = iota
	TokenEmphasis
	TokenStrong
	TokenCode
)

type InlineToken struct {
	Type InlineTokenType
	Text string
	URL  string
}

func TokenizeInline(input string) []InlineToken {
	var tokens []InlineToken

	for len(input) > 0 {
		switch {
		case strings.HasPrefix(input, "**"):
			end := strings.Index(input[2:], "**")
			if end >= 0 {
				text := input[2 : 2+end]
				tokens = append(tokens, InlineToken{Type: TokenStrong, Text: text})
				input = input[2+end+2:]
				continue
			}
		case strings.HasPrefix(input, "*"):
			end := strings.Index(input[1:], "*")
			if end >= 0 {
				text := input[1 : 1+end]
				tokens = append(tokens, InlineToken{Type: TokenEmphasis, Text: text})
				input = input[1+end+1:]
				continue
			}
		case strings.HasPrefix(input, "`"):
			end := strings.Index(input[1:], "`")
			if end >= 0 {
				text := input[1 : 1+end]
				tokens = append(tokens, InlineToken{Type: TokenCode, Text: text})
				input = input[1+end+1:]
				continue
			}
		}

		tokens = append(tokens, InlineToken{Type: TokenText, Text: string(input[0])})
		input = input[1:]
	}

	return mergeTextTokens(tokens)
}

func mergeTextTokens(tokens []InlineToken) []InlineToken {
	if len(tokens) == 0 {
		return nil
	}

	var merged []InlineToken
	current := tokens[0]

	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		if current.Type == TokenText && t.Type == TokenText {
			current.Text += t.Text
			continue
		}

		merged = append(merged, current)
		current = t
	}

	merged = append(merged, current)
	return merged
}

func RenderInline(tokens []InlineToken) string {
	var b strings.Builder

	for _, token := range tokens {
		switch token.Type {
		case TokenText:
			b.WriteString(token.Text)
		case TokenEmphasis:
			b.WriteString("<em>" + token.Text + "</em>")
		case TokenStrong:
			b.WriteString("<strong>" + token.Text + "</strong>")
		case TokenCode:
			b.WriteString("<code>" + token.Text + "</code>")
		}
	}

	return b.String()
}

func renderParagraph(line string) string {
	return "<p>" + RenderInline(TokenizeInline(line)) + "</p>"
}
