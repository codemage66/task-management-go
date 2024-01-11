package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDatabase struct {
	Database
}

var (
	db          Database
	testConnStr string = "postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable"
)

func TestRealDatabase(t *testing.T) {
	db = NewDatabase()
	assert.NotNil(t, db)
}

func TestConnect(t *testing.T) {
	conn, err := db.Connect(testConnStr)
	assert.NoError(t, err, "Test Connect should not return an error")
	assert.NotNil(t, conn, "Test Connect should return a connection")
	assert.NoError(t, conn.Ping(), "Test Connect should return a valid connection")
	assert.NoError(t, conn.Close(), "Test Connect should close the connection")
	assert.Error(t, conn.Ping(), "Test Connect should close the connection")

	defer conn.Close()
}

func TestRunMigration(t *testing.T) {
	err := db.RunMigration(testConnStr, "file:///migration")

	assert.NoError(t, err, "Test RunMigration should not return an error")
	assert.Nil(t, err, "Test RunMigration should set the migration to nil")
}

func TestRollbackMigration(t *testing.T) {
	err := db.RollbackMigration(testConnStr, "file:///migration")

	assert.NoError(t, err, "Test RollbackMigration should not return an error")
	assert.Nil(t, err, "Test RollbackMigration should set the migration to nil")
}

func TestRollbackMigrationError(t *testing.T) {
	err := db.RollbackMigration("bad connection string", "file:///migration")

	assert.Error(t, err, "Test RollbackMigration should return an error")
	assert.NotNil(t, err, "Test RollbackMigration should not set the migration to nil")
}

func TestConnectNil(t *testing.T) {
	conn, err := db.Connect("")

	assert.Error(t, err, "Test Connect should return an error")
	assert.Nil(t, conn, "Test Connect should not return a connection")
}

func TestConnectError(t *testing.T) {
	conn, err := db.Connect("bad connection string")

	assert.Error(t, err, "Test Connect should return an error")
	assert.Nil(t, conn, "Test Connect should not return a connection")
}

func TestRunMigrationError(t *testing.T) {
	err := db.RunMigration("bad connection string", "file:///migration")

	assert.Error(t, err, "Test RunMigration should return an error")
	assert.NotNil(t, err, "Test RunMigration should not set the migration to nil")
}
