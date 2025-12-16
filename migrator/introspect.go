package migrator

import "database/sql"

type Column struct {
	Name     string
	Type     string
	Nullable bool
}

func LoadSchema(db *sql.DB) (map[string][]Column, error) {
	rows, err := db.Query(`
		SELECT table_name, column_name, data_type, is_nullable, is_nullable
		FROM information_schema.columns
		WHERE table_schema = 'public'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schema := map[string][]Column{}

	for rows.Next() {
		var table, col, typ, nullable string
		rows.Scan(&table, &col, &typ, &nullable)

		schema[table] = append(schema[table], Column{
			Name:     col,
			Type:     typ,
			Nullable: nullable == "YES",
		})
	}
	return schema, nil
}
