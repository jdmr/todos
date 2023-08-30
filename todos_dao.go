package main

import (
	"database/sql"
	"errors"
)

type TodoDao interface {
	GetAll() ([]*Todo, error)
	Get(id string) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id string) error
}

type TodoDaoImpl struct {
	conn *sql.DB
}

func NewTodoDao(conn *sql.DB) TodoDao {
	return &TodoDaoImpl{conn: conn}
}

func (dao *TodoDaoImpl) GetAll() ([]*Todo, error) {
	return nil, errors.New("not implemented yet")
}

func (dao *TodoDaoImpl) Get(id string) (*Todo, error) {
	return nil, errors.New("not implemented yet")
}

func (dao *TodoDaoImpl) Create(todo *Todo) (*Todo, error) {
	return nil, errors.New("not implemented yet")
}

func (dao *TodoDaoImpl) Update(todo *Todo) (*Todo, error) {
	return nil, errors.New("not implemented yet")
}

func (dao *TodoDaoImpl) Delete(id string) error {
	return errors.New("not implemented yet")
}
