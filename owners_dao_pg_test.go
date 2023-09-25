package main

import (
	"database/sql"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestPGGetAllOwners(t *testing.T) {
	t.Log("Testing Postgres GetAllOwners")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")

	dao := NewPGOwnerDao(conn)
	owners, err := dao.GetAll()
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	if len(owners) == 0 {
		cleanupOwners(conn)
		t.Fatal("expected at least one owner")
	}

	cleanupOwners(conn)
}

func TestPGGetOwner(t *testing.T) {
	t.Log("Testing Postgres GetOwner")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")

	dao := NewPGOwnerDao(conn)
	owner, err := dao.Get("test")
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	if owner == nil {
		cleanupOwners(conn)
		t.Fatal("expected an owner")
	}

	cleanupOwners(conn)
}

func TestPGCreateOwner(t *testing.T) {
	t.Log("Testing Postgres CreateOwner")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dao := NewPGOwnerDao(conn)
	owner := &Owner{
		ID:   "test",
		Name: "test",
	}
	err = dao.Create(owner)
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	cleanupOwners(conn)
}

func TestPGUpdateOwner(t *testing.T) {
	t.Log("Testing Postgres UpdateOwner")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")

	dao := NewPGOwnerDao(conn)
	owner := &Owner{
		ID:   "test",
		Name: "testing",
	}
	err = dao.Update(owner)
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	if owner.Name != "testing" {
		cleanupOwners(conn)
		t.Fatal("expected owner name to be 'testing'")
	}

	cleanupOwners(conn)
}

func TestPGDeleteOwner(t *testing.T) {
	t.Log("Testing Postgres DeleteOwner")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")

	dao := NewPGOwnerDao(conn)
	err = dao.Delete("test")
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	var cont int
	err = conn.QueryRow("select count(*) from owners where id = 'test'").Scan(&cont)
	if err != nil {
		cleanupOwners(conn)
		t.Fatal(err)
	}

	if cont != 0 {
		cleanupOwners(conn)
		t.Fatal("expected owner to be deleted")
	}

	cleanupOwners(conn)
}

func cleanupOwners(conn *sql.DB) {
	conn.Exec("delete from owners where id like 'test%'")
}
