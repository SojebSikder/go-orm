package parser

import "fmt"

type ParserState struct {
	toks []Token
	i    int
}

func (p *ParserState) peek() *Token {
	if p.i >= len(p.toks) {
		return nil
	}
	return &p.toks[p.i]
}

func (p *ParserState) next() *Token {
	if p.i >= len(p.toks) {
		return nil
	}
	t := &p.toks[p.i]
	p.i++
	return t
}

func (p *ParserState) expect(val string) (*Token, error) {
	tok := p.next()
	if tok == nil || tok.Val != val {
		return nil, fmt.Errorf("expected '%s', got '%v'", val, tok)
	}
	return tok, nil
}

func (p *ParserState) expectType(tt TokenType) (*Token, error) {
	tok := p.next()
	if tok == nil || tok.Typ != tt {
		return nil, fmt.Errorf("expected token type %v, got %v", tt, tok)
	}
	return tok, nil
}
