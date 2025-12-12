package parser

import "strings"

func collectUntilRBrace(ps *ParserState) []Token {
	var out []Token
	depth := 0
	if ps.peek() != nil && ps.peek().Typ == T_LBRACE {
		ps.next()
		depth = 1
	}
	for ps.peek() != nil && depth > 0 {
		t := ps.next()
		if t.Typ == T_LBRACE {
			depth++
		} else if t.Typ == T_RBRACE {
			depth--
			if depth == 0 {
				break
			}
		}
		if depth > 0 {
			out = append(out, *t)
		}
	}
	return out
}

func parseModel(ps *ParserState) (*Model, error) {
	ps.next() // consume 'model'
	nameTok := ps.next()
	name := ""
	if nameTok != nil {
		name = nameTok.Val
	}

	toks := collectUntilRBrace(ps)
	m := &Model{Name: name}
	i := 0

	for i < len(toks) {
		tok := toks[i]
		if tok.Typ == T_DAT {
			attrParts := []string{tok.Val}
			j := i + 1
			for j < len(toks) && toks[j].Typ != T_IDENT && toks[j].Typ != T_AT && toks[j].Typ != T_LBRACE && toks[j].Typ != T_RBRACE {
				attrParts = append(attrParts, toks[j].Val)
				j++
			}
			if j < len(toks) {
				attrParts = append(attrParts, toks[j].Val)
				j++
			}
			m.Attributes = append(m.Attributes, strings.Join(attrParts, " "))
			i = j
			continue
		}

		if tok.Typ != T_IDENT {
			i++
			continue
		}

		fieldName := tok.Val
		i++
		if i >= len(toks) {
			break
		}
		typTok := toks[i]
		typ := typTok.Val
		isArray := false
		isOptional := false
		if strings.HasSuffix(typ, "[]") {
			isArray = true
			typ = strings.TrimSuffix(typ, "[]")
		}
		if strings.HasSuffix(typ, "?") {
			isOptional = true
			typ = strings.TrimSuffix(typ, "?")
		}

		if i+1 < len(toks) && toks[i+1].Typ == T_LBRACK && i+2 < len(toks) && toks[i+2].Typ == T_RBRACK {
			isArray = true
			i += 2
		}

		attrs := []string{}
		rawParts := []string{fieldName, typTok.Val}
		j := i + 1
		for j < len(toks) {
			if toks[j].Typ == T_AT {
				attrs = append(attrs, toks[j].Val)
				rawParts = append(rawParts, toks[j].Val)
				j++
				continue
			}
			if toks[j].Typ == T_IDENT {
				break
			}
			rawParts = append(rawParts, toks[j].Val)
			j++
		}

		field := Field{
			Name:       fieldName,
			Type:       typ,
			IsArray:    isArray,
			IsOptional: isOptional,
			Attributes: attrs,
			Raw:        strings.Join(rawParts, " "),
		}
		m.Fields = append(m.Fields, field)
		i = j
	}
	return m, nil
}
