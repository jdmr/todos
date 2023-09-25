package main

type TodoDao interface {
	GetAll() ([]*Todo, error)
	Get(id string) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id string) error
	Done(id string) error
}
