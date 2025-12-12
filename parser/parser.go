package parser

func ParseSchema(src string) (*SchemaAST, error) {
	clean := StripComments(src)
	toks := Tokenize(clean)
	ps := &ParserState{toks: toks, i: 0}
	ast := &SchemaAST{Raw: src}

	for ps.peek() != nil {
		t := ps.peek()
		switch {
		case t.Typ == T_IDENT && t.Val == "generator":
			gen, err := parseGenerator(ps)
			if err != nil {
				return nil, err
			}
			ast.Generators = append(ast.Generators, *gen)
		case t.Typ == T_IDENT && t.Val == "datasource":
			ds, err := parseDatasource(ps)
			if err != nil {
				return nil, err
			}
			ast.Datasources = append(ast.Datasources, *ds)
		case t.Typ == T_IDENT && t.Val == "model":
			m, err := parseModel(ps)
			if err != nil {
				return nil, err
			}
			ast.Models = append(ast.Models, *m)
		case t.Typ == T_IDENT && t.Val == "enum":
			e, err := parseEnum(ps)
			if err != nil {
				return nil, err
			}
			ast.Enums = append(ast.Enums, *e)
		default:
			ps.next()
		}
	}
	return ast, nil
}
