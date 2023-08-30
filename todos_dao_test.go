package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestGetAll(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://jdmr@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dao := NewTodoDao(conn)
	todos, err := dao.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) == 0 {
		t.Fatal("expected at least one todo")
	}
}
