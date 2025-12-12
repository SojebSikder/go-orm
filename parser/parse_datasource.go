package parser

func parseDatasource(ps *ParserState) (*Datasource, error) {
	ps.next()
	nameTok := ps.next()
	name := ""
	if nameTok != nil {
		name = nameTok.Val
	}

	for ps.peek() != nil && ps.peek().Typ != T_LBRACE {
		ps.next()
	}
	if ps.peek() == nil {
		return &Datasource{Name: name}, nil
	}
	ps.next() // consume {

	fields := map[string]string{}
	for ps.peek() != nil && ps.peek().Typ != T_RBRACE {
		identTok := ps.next()
		if identTok == nil || identTok.Typ != T_IDENT {
			continue
		}
		for ps.peek() != nil && ps.peek().Typ != T_EQUAL {
			ps.next()
		}
		if ps.peek() != nil && ps.peek().Typ == T_EQUAL {
			ps.next()
		}
		valTok := ps.next()
		val := ""
		if valTok != nil {
			val = valTok.Val
		}
		fields[identTok.Val] = val
	}

	if ps.peek() != nil && ps.peek().Typ == T_RBRACE {
		ps.next()
	}
	return &Datasource{Name: name, Fields: fields}, nil
}
