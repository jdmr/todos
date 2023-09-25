package main

type OwnerDao interface {
	GetAll() ([]*Owner, error)
	Get(id string) (*Owner, error)
	Create(owner *Owner) error
	Update(owner *Owner) error
	Delete(id string) error
}
