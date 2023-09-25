package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TodoDaoPGImpl struct {
	conn *sql.DB
}

func NewPGTodoDao(conn *sql.DB) TodoDao {
	return &TodoDaoPGImpl{conn: conn}
}

func (dao *TodoDaoPGImpl) GetAll() ([]*Todo, error) {
	rows, err := dao.conn.Query(`
		SELECT 
			t.id
			, t.title
			, t.completed
			, t.created_at
			, t.updated_at
			, o.id
			, o.name
		FROM todos t
		JOIN owners o ON t.owner_id = o.id
		order by t.created_at desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		todo := &Todo{
			Owner: &Owner{},
		}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.Owner.ID, &todo.Owner.Name)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (dao *TodoDaoPGImpl) Get(id string) (*Todo, error) {
	todo := &Todo{
		Owner: &Owner{},
	}
	err := dao.conn.QueryRow(`
		SELECT 
			t.id
			, t.title
			, t.completed
			, t.created_at
			, t.updated_at
			, o.id
			, o.name
		FROM todos t
		JOIN owners o ON t.owner_id = o.id
		WHERE t.id = $1
	`, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.Owner.ID, &todo.Owner.Name)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *TodoDaoPGImpl) Create(todo *Todo) error {
	_, err := dao.conn.Exec("INSERT INTO todos (id, title, completed, created_at, updated_at, owner_id) VALUES ($1, $2, $3, NOW(), NOW(), $4)", todo.ID, todo.Title, todo.Completed, todo.Owner.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoPGImpl) Update(todo *Todo) error {
	_, err := dao.conn.Exec("UPDATE todos SET title = $1, completed = $2, updated_at = now() WHERE id = $3", todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoPGImpl) Delete(id string) error {
	_, err := dao.conn.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoPGImpl) Done(id string) error {
	_, err := dao.conn.Exec("UPDATE todos SET completed = true, updated_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
