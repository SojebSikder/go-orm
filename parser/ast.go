package parser

type SchemaAST struct {
	Generators  []Generator  `json:"generators,omitempty"`
	Datasources []Datasource `json:"datasources,omitempty"`
	Models      []Model      `json:"models,omitempty"`
	Enums       []Enum       `json:"enums,omitempty"`
	Raw         string       `json:"-"`
}

type Generator struct {
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields,omitempty"`
}

type Datasource struct {
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields,omitempty"`
}

type Model struct {
	Name        string            `json:"name"`
	Fields      []Field           `json:"fields,omitempty"`
	Attributes  []string          `json:"attributes,omitempty"`
	RawContents []string          `json:"raw_contents,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

type Field struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	IsArray    bool     `json:"is_array"`
	IsOptional bool     `json:"is_optional"`
	Attributes []string `json:"attributes,omitempty"`
	Raw        string   `json:"raw,omitempty"`
}

type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values,omitempty"`
}
