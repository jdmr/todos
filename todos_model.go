package main

import "time"

type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Owner     *Owner    `json:"owner"`
}
