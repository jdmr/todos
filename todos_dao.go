package main

import (
	"database/sql"
)

type TodoDao interface {
	GetAll() ([]*Todo, error)
	Get(id string) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id string) error
}

type TodoDaoImpl struct {
	conn *sql.DB
}

func NewTodoDao(conn *sql.DB) TodoDao {
	return &TodoDaoImpl{conn: conn}
}

func (dao *TodoDaoImpl) GetAll() ([]*Todo, error) {
	rows, err := dao.conn.Query("SELECT id, title, completed, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (dao *TodoDaoImpl) Get(id string) (*Todo, error) {
	todo := &Todo{}
	err := dao.conn.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *TodoDaoImpl) Create(todo *Todo) error {
	_, err := dao.conn.Exec("INSERT INTO todos (id, title, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", todo.ID, todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoImpl) Update(todo *Todo) error {
	_, err := dao.conn.Exec("UPDATE todos SET title = $1, completed = $2, updated_at = now() WHERE id = $3", todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoImpl) Delete(id string) error {
	_, err := dao.conn.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
