package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sojebsikder/go-orm/parser"
)

type PostgreSQLGenerator struct {
	schema *parser.SchemaAST
}

func NewPostgreSQLGenerator(s *parser.SchemaAST) *PostgreSQLGenerator {
	return &PostgreSQLGenerator{schema: s}
}

func (g *PostgreSQLGenerator) Generate() string {
	var out strings.Builder

	out.WriteString("-- Auto-generated PostgreSQL schema\n\n")
	// out.WriteString("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\";\n\n")

	// Enums
	for _, e := range g.schema.Enums {
		out.WriteString(g.generateEnum(e))
		out.WriteString("\n")
	}

	// Tables
	for _, m := range g.schema.Models {
		out.WriteString(g.generateModel(m))
		out.WriteString("\n")
	}

	return out.String()
}

func (g *PostgreSQLGenerator) generateEnum(e parser.Enum) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		"CREATE TYPE %s AS ENUM (\n",
		snake(e.Name),
	))

	values := []string{}
	for _, v := range e.Values {
		values = append(values, fmt.Sprintf("  '%s'", v))
	}

	sb.WriteString(strings.Join(values, ",\n"))
	sb.WriteString("\n);\n")

	return sb.String()
}

func (g *PostgreSQLGenerator) generateModel(m parser.Model) string {
	var sb strings.Builder
	table := resolveTableName(m)

	sb.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", table))

	cols := []string{}
	constraints := []string{}

	for _, f := range m.Fields {
		if isMetaField(f) {
			continue
		}

		if isRelationField(f) {
			constraints = append(constraints, parseRelationConstraint(f, table))
			continue
		}

		cols = append(cols, "  "+renderColumn(f))

		if isPrimaryKey(f) {
			constraints = append(constraints, fmt.Sprintf("  PRIMARY KEY (%s)", snake(f.Name)))
		}
	}

	sb.WriteString(strings.Join(cols, ",\n"))

	if len(constraints) > 0 {
		sb.WriteString(",\n")
		sb.WriteString(strings.Join(uniqueStrings(constraints), ",\n"))
	}

	sb.WriteString("\n);\n")
	return sb.String()
}

// Column Rendering

func renderColumn(f parser.Field) string {
	colType := mapType(f)
	nullable := "NOT NULL"

	if strings.Contains(f.Raw, "?") || f.IsOptional {
		nullable = "NULL"
	}

	def := extractDefault(f)

	parts := []string{snake(f.Name), colType, nullable}
	if def != "" {
		parts = append(parts, "DEFAULT "+def)
	}

	return strings.Join(parts, " ")
}

// Type Mapping

func mapType(f parser.Field) string {
	if hasAttr(f, "@db.Text") {
		return "TEXT"
	}
	if hasAttr(f, "@db.SmallInt") {
		return "SMALLINT"
	}

	switch f.Type {
	case "String":
		return "TEXT"
	case "DateTime":
		return "TIMESTAMPTZ"
	case "Int":
		return "INTEGER"
	case "Boolean":
		return "BOOLEAN"
	default:
		return "TEXT"
	}
}

// Defaults

func extractDefault(f parser.Field) string {
	if strings.Contains(f.Raw, "now()") {
		return "now()"
	}
	if strings.Contains(f.Raw, "cuid()") {
		return "gen_random_uuid()"
	}

	re := regexp.MustCompile(`@default\\((.*?)\\)`)
	m := re.FindStringSubmatch(f.Raw)
	if len(m) == 2 {
		if v, err := strconv.Atoi(m[1]); err == nil {
			return strconv.Itoa(v)
		}
	}

	return ""
}

// Relations

func parseRelationConstraint(f parser.Field, table string) string {
	for _, a := range f.Attributes {
		if strings.HasPrefix(a, "@relation") {
			re := regexp.MustCompile(`fields:\s*\[(.*?)\].*references:\s*\[(.*?)\].*onDelete:\s*(\w+)`)
			m := re.FindStringSubmatch(a)
			if len(m) == 4 {
				col := snake(m[1])
				refTable := snakePlural(f.Type)
				refCol := snake(m[2])
				return fmt.Sprintf("  FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE %s", col, refTable, refCol, strings.ToUpper(m[3]))
			}
		}
	}
	return ""
}

// Helpers

func isPrimaryKey(f parser.Field) bool {
	return hasAttr(f, "@id")
}

func hasAttr(f parser.Field, attr string) bool {
	for _, a := range f.Attributes {
		if strings.Contains(a, attr) {
			return true
		}
	}
	return false
}

func isRelationField(f parser.Field) bool {
	return f.Type != "String" && f.Type != "DateTime" && f.Type != "Int" && f.Type != "Boolean"
}

func isMetaField(f parser.Field) bool {
	return strings.HasPrefix(strings.TrimSpace(f.Raw), "@@") || strings.HasPrefix(f.Name, "map")
}

func resolveTableName(m parser.Model) string {
	// for _, f := range m.Fields {
	// 	if strings.HasPrefix(f.Raw, "map") {
	// 		return extractMapName(f.Raw)
	// 	}
	// }
	return snakePlural(m.Name)
}

func extractMapName(raw string) string {
	re := regexp.MustCompile(`map\\s*\\(\\s*(\\w+)\\s*\\)`)
	m := re.FindStringSubmatch(raw)
	if len(m) == 2 {
		return m[1]
	}
	return ""
}

func snake(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(s, "${1}_${2}"))
}

func snakePlural(s string) string {
	return snake(s) + "s"
}

func uniqueStrings(in []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, v := range in {
		if v != "" && !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}
