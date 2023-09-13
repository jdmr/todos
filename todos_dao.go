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
	Done(id string) error
	GetOwners() ([]*Owner, error)
}

type TodoDaoImpl struct {
	conn *sql.DB
}

func NewTodoDao(conn *sql.DB) TodoDao {
	return &TodoDaoImpl{conn: conn}
}

func (dao *TodoDaoImpl) GetAll() ([]*Todo, error) {
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

func (dao *TodoDaoImpl) Get(id string) (*Todo, error) {
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

func (dao *TodoDaoImpl) Create(todo *Todo) error {
	_, err := dao.conn.Exec("INSERT INTO todos (id, title, completed, created_at, updated_at, owner_id) VALUES ($1, $2, $3, NOW(), NOW(), $4)", todo.ID, todo.Title, todo.Completed, todo.Owner.ID)
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

func (dao *TodoDaoImpl) Done(id string) error {
	_, err := dao.conn.Exec("UPDATE todos SET completed = true, updated_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDaoImpl) GetOwners() ([]*Owner, error) {
	rows, err := dao.conn.Query("SELECT id, name FROM owners")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	owners := []*Owner{}
	for rows.Next() {
		owner := &Owner{}
		err := rows.Scan(&owner.ID, &owner.Name)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}
	return owners, nil
}
