package parser

func parseEnum(ps *ParserState) (*Enum, error) {
	ps.next() // consume 'enum'
	nameTok := ps.next()
	name := ""
	if nameTok != nil {
		name = nameTok.Val
	}

	toks := collectUntilRBrace(ps)
	e := &Enum{Name: name}
	for _, t := range toks {
		if t.Typ == T_IDENT {
			e.Values = append(e.Values, t.Val)
		}
	}
	return e, nil
}
