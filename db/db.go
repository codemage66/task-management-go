package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

// Database is an interface for database
type Database interface {
	Connect(string) (*sql.DB, error)
	RunMigration(string, string) error
	RollbackMigration(string, string) error
}

// DatabaseInstance is a real implementation of Database
type DatabaseInstance struct{}

// NewDatabase returns a new instance of Database
func NewDatabase() *DatabaseInstance {
	return &DatabaseInstance{}
}

// Connect to database using connection string
func (rdb *DatabaseInstance) Connect(connStr string) (*sql.DB, error) {
	const maxOpenConns = 100
	const maxIdleConns = 10

	connector, err := pq.NewConnector(connStr)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Minute * 1)

	// Check if database is alive
	if err = db.PingContext(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}

// RunMigration runs migration
func (rdb *DatabaseInstance) RunMigration(connStr, migrateDir string) error {
	db, err := rdb.Connect(connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	migration, err := migrate.New(migrateDir, connStr)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil {

		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	defer migration.Close()
	return nil
}

func (rdb *DatabaseInstance) RollbackMigration(connStr, migrateDir string) error {
	migration, err := migrate.New(migrateDir, connStr)
	if err != nil {
		return err
	}

	if err := migration.Down(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return err
	}

	defer migration.Close()
	return nil
}
