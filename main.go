package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}
	defer conn.Close()

	dao := NewPGTodoDao(conn)
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Welcome to Todo CLI")
	log.Println("Enter command:")
	scanner.Scan()
	cmd := scanner.Text()
	for cmd != "quit" {
		switch cmd {
		case "list":
			log.Println("list of todos")
			todos, err := dao.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			for _, todo := range todos {
				log.Println("********************************")
				log.Printf("%s\t%s\t%t\t%s(%s)\n", todo.ID, todo.Title, todo.Completed, todo.Owner.Name, todo.Owner.ID)
			}
		case "create":
			log.Println("Enter todo title:")
			scanner.Scan()
			title := scanner.Text()
			id := uuid.New().String()
			id = id[:3]
			todo := &Todo{ID: id, Title: title, Owner: &Owner{}}
			log.Println("List of owners:")
			owners, err := dao.GetOwners()
			if err != nil {
				log.Fatal(err)
			}
			for _, owner := range owners {
				log.Printf("%s\t%s\n", owner.ID, owner.Name)
			}
			log.Print("Enter todo owner id: ")
			scanner.Scan()
			todo.Owner.ID = scanner.Text()

			err = dao.Create(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo created")
		case "update":
			log.Println("Enter todo id:")
			scanner.Scan()
			id := scanner.Text()
			log.Println("Enter todo title:")
			scanner.Scan()
			title := scanner.Text()
			log.Println("Enter todo completed:")
			scanner.Scan()
			completed := scanner.Text()
			todo := &Todo{ID: id, Title: title, Completed: completed == "true"}
			err := dao.Update(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo updated")
		case "delete":
			log.Println("Enter todo id:")
			scanner.Scan()
			id := scanner.Text()
			err := dao.Delete(id)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo deleted")
		case "done":
			log.Println("Enter todo id:")
			scanner.Scan()
			id := scanner.Text()
			err := dao.Done(id)
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("Unknown command")
		}
		log.Println("Enter command:")
		scanner.Scan()
		cmd = scanner.Text()
	}
	log.Println("Bye!")
}
