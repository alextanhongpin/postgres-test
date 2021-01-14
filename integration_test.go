// +build integration

package main_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alextanhongpin/postgres"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Sum(ctx context.Context) (int, error) {
	var i int
	err := r.db.QueryRowContext(ctx, `SELECT 1 + 1`).Scan(&i)
	return i, err
}

func TestMain(m *testing.M) {
	_, cleanup := postgres.SetupTestDB()
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	db, err := postgres.NewTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepository(db)
	sum, err := repo.Sum(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	expected := 2
	if got := sum; sum != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
