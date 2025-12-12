package parser

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
