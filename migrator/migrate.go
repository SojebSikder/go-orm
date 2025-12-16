package migrator

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Migrator struct {
	db            *sql.DB
	migrationsDir string
}

func NewMigrator(db *sql.DB, dir string) *Migrator {
	return &Migrator{
		db:            db,
		migrationsDir: dir,
	}
}

func (m *Migrator) EnsureMigrationTable() error {
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS _orm_migrations (
		  id SERIAL PRIMARY KEY,
		  version_id TEXT NOT NULL,
		  is_applied BOOLEAN NOT NULL DEFAULT FALSE,
		  checksum TEXT NOT NULL,
		  timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
		)
	`)

	return err
}

func (m *Migrator) Apply() error {
	if err := m.EnsureMigrationTable(); err != nil {
		return err
	}

	applied, err := m.appliedMigrations()
	if err != nil {
		return err
	}

	return filepath.WalkDir(m.migrationsDir, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() || filepath.Ext(path) != ".sql" {
			return nil
		}

		name := filepath.Base(path)
		if applied[name] {
			return nil
		}

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		checksum := hash(sqlBytes)

		tx, err := m.db.Begin()
		if err != nil {
			return err
		}

		if _, err := tx.Exec(string(sqlBytes)); err != nil {
			tx.Rollback()
			return fmt.Errorf("migrations %s failed: %w", name, err)
		}

		if _, err := tx.Exec(
			`INSERT INTO _orm_migrations (version_id, checksum) VALUES ($1,$2)`,
			name, checksum,
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert migration %s: %w", name, err)
		}

		return tx.Commit()
	})
}

func (m *Migrator) appliedMigrations() (map[string]bool, error) {
	rows, err := m.db.Query(`SELECT version_id FROM _orm_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[string]bool{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		out[name] = true
	}
	return out, nil
}

func hash(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}
