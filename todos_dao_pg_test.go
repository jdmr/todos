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

func TestPGGetAll(t *testing.T) {
	t.Log("Testing Postgres GetAll")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")
	conn.Exec("insert into todos (id, title, completed, created_at, updated_at, owner_id) values ('test', 'test', false, now(), now(), 'test')")

	dao := NewPGTodoDao(conn)
	todos, err := dao.GetAll()
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	if len(todos) == 0 {
		cleanup(conn)
		t.Fatal("expected at least one todo")
	}

	cleanup(conn)
}

func TestPGGet(t *testing.T) {
	t.Log("Testing Postgres Get")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")
	conn.Exec("insert into todos (id, title, completed, created_at, updated_at, owner_id) values ('test', 'test', false, now(), now(), 'test')")

	dao := NewPGTodoDao(conn)
	todo, err := dao.Get("test")
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	if todo == nil {
		cleanup(conn)
		t.Fatal("expected a todo")
	}

	if todo.Title != "test" {
		cleanup(conn)
		t.Fatal("expected todo title to be 'test'")
	}

	cleanup(conn)
}

func TestPGCreate(t *testing.T) {
	t.Logf("Testing Postgres Create")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")

	dao := NewPGTodoDao(conn)
	todo := &Todo{ID: "test", Title: "test", Completed: false, Owner: &Owner{ID: "test"}}
	err = dao.Create(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func TestPGUpdate(t *testing.T) {
	t.Logf("Testing Postgres Update")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")
	conn.Exec("insert into todos (id, title, completed, created_at, updated_at, owner_id) values ('test', 'test', false, now(), now(), 'test')")

	dao := NewPGTodoDao(conn)
	todo := &Todo{ID: "test", Title: "test", Completed: true}
	err = dao.Update(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func TestPGDelete(t *testing.T) {
	t.Logf("Testing Postgres Delete")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into owners (id, name) values ('test', 'test')")
	conn.Exec("insert into todos (id, title, completed, created_at, updated_at, owner_id) values ('test', 'test', false, now(), now(), 'test')")

	dao := NewPGTodoDao(conn)
	err = dao.Delete("test")
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func cleanup(conn *sql.DB) {
	conn.Exec("delete from owners where name like 'test%'")
}
