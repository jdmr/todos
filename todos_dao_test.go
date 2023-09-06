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
	t.Log("Testing GetAll")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
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

func TestGet(t *testing.T) {
	t.Log("Testing Get")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.QueryRow("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
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

func TestCreate(t *testing.T) {
	t.Logf("Testing Create")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dao := NewTodoDao(conn)
	todo := &Todo{ID: "test", Title: "test"}
	err = dao.Create(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func TestUpdate(t *testing.T) {
	t.Logf("Testing Update")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
	todo := &Todo{ID: "test", Title: "test", Completed: true}
	err = dao.Update(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func TestDelete(t *testing.T) {
	t.Logf("Testing Delete")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
	err = dao.Delete("test")
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func cleanup(conn *sql.DB) {
	conn.Exec("delete from todos where title like 'test%'")
}
