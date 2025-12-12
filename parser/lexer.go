package parser

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	_ TokenType = iota
	T_IDENT
	T_LBRACE
	T_RBRACE
	T_LPAREN
	T_RPAREN
	T_AT
	T_DAT
	T_LBRACK
	T_RBRACK
	T_SEMI
	T_EQUAL
	T_COMMA
	T_STRING
	T_OTHER
)

type Token struct {
	Typ TokenType
	Val string
}

func isIdentStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_' || r == '@' || r == '$'
}

func isIdentPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' || r == '-'
}

func Tokenize(input string) []Token {
	var toks []Token
	i := 0
	n := len(input)

	for i < n {
		c := input[i]
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			i++
			continue
		}

		switch c {
		case '{':
			toks = append(toks, Token{T_LBRACE, "{"})
			i++
			continue
		case '}':
			toks = append(toks, Token{T_RBRACE, "}"})
			i++
			continue
		case '(':
			toks = append(toks, Token{T_LPAREN, "("})
			i++
			continue
		case ')':
			toks = append(toks, Token{T_RPAREN, ")"})
			i++
			continue
		case '[':
			toks = append(toks, Token{T_LBRACK, "["})
			i++
			continue
		case ']':
			toks = append(toks, Token{T_RBRACK, "]"})
			i++
			continue
		case ';':
			toks = append(toks, Token{T_SEMI, ";"})
			i++
			continue
		case ',':
			toks = append(toks, Token{T_COMMA, ","})
			i++
			continue
		case '=':
			toks = append(toks, Token{T_EQUAL, "="})
			i++
			continue
		case '"', '\'':
			quote := c
			j := i + 1
			var sb strings.Builder
			for j < n {
				if input[j] == '\\' && j+1 < n {
					sb.WriteByte(input[j])
					j++
					sb.WriteByte(input[j])
					j++
					continue
				}
				if input[j] == quote {
					j++
					break
				}
				sb.WriteByte(input[j])
				j++
			}
			toks = append(toks, Token{T_STRING, sb.String()})
			i = j
			continue
		case '@':
			if i+1 < n && input[i+1] == '@' {
				toks = append(toks, Token{T_DAT, "@@"})
				i += 2
				continue
			}
			j := i
			var sb strings.Builder
			sb.WriteByte(input[j])
			j++
			for j < n {
				r := rune(input[j])
				if r == '(' {
					parenCount := 0
					for j < n {
						sb.WriteByte(input[j])
						if input[j] == '(' {
							parenCount++
						} else if input[j] == ')' {
							parenCount--
							if parenCount == 0 {
								j++
								break
							}
						}
						j++
					}
					break
				}
				if isIdentPart(r) {
					sb.WriteByte(input[j])
					j++
					continue
				}
				break
			}
			toks = append(toks, Token{T_AT, sb.String()})
			i = j
			continue
		}

		if isIdentStart(rune(c)) || unicode.IsDigit(rune(c)) {
			j := i
			for j < n && (isIdentPart(rune(input[j])) || input[j] == '@') {
				j++
			}
			toks = append(toks, Token{T_IDENT, input[i:j]})
			i = j
			continue
		}

		toks = append(toks, Token{T_OTHER, string(c)})
		i++
	}

	return toks
}
