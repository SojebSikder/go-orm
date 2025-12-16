package migrator

import "fmt"

func Diff(current map[string][]Column, desired map[string][]Column) []string {
	var sql []string

	for table, cols := range desired {
		if _, ok := current[table]; !ok {
			// new table
			continue
		}

		existing := map[string]Column{}
		for _, c := range current[table] {
			existing[c.Name] = c
		}

		for _, c := range cols {
			if _, ok := existing[c.Name]; !ok {
				sql = append(sql,
					fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s",
						table, c.Name, c.Type),
				)
			}
		}
	}

	return sql
}
