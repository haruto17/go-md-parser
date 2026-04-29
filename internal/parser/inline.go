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
