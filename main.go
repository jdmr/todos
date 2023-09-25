package main

import (
	"bufio"
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		log.Fatal("Could not open mongodb connection: ", err)
	}

	var todoDao TodoDao
	var ownerDao OwnerDao
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("********************************")
	log.Println("Welcome to Todo CLI")
	log.Println("********************************")
	log.Println("Commands:")
	log.Println("pg - select postgres database")
	log.Println("mg - select mongodb database")
	log.Println("list - list all todos")
	log.Println("create - create a todo")
	log.Println("update - update a todo")
	log.Println("delete - delete a todo")
	log.Println("done - mark a todo as done")
	log.Println("quit - quit the program")
	log.Println("********************************")
	log.Print("Enter command: ")
	scanner.Scan()
	cmd := scanner.Text()
	for cmd != "quit" {
		switch cmd {
		case "pg":
			todoDao = NewPGTodoDao(conn)
			ownerDao = NewPGOwnerDao(conn)
		case "mg":
			todoDao = NewMongoTodoDao(client)
			ownerDao = NewMongoOwnerDao(client)
		case "list":
			if todoDao == nil {
				log.Println("Please select a database first")
				break
			}
			log.Println("list of todos")
			todos, err := todoDao.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			if len(todos) == 0 {
				log.Println("********************************")
				log.Println("********************************")
				log.Println("No todos found")
				log.Println("********************************")
				break
			}
			for _, todo := range todos {
				log.Println("********************************")
				// log.Printf("ID: %s\n", todo.ID)
				log.Printf("%s\t%s\t%t\t%s(%s)\n", todo.ID, todo.Title, todo.Completed, todo.Owner.Name, todo.Owner.ID)
			}
		case "create":
			if todoDao == nil {
				log.Println("Please select a database first")
				break
			}
			log.Print("Enter todo title: ")
			scanner.Scan()
			title := scanner.Text()
			id := uuid.New().String()
			id = id[:3]
			todo := &Todo{ID: id, Title: title, Owner: &Owner{}}
			log.Println("List of owners:")
			owners, err := ownerDao.GetAll()
			if err != nil {
				log.Fatal(err)
			}
			for _, owner := range owners {
				log.Printf("%s\t%s\n", owner.ID, owner.Name)
			}
			log.Print("Enter todo owner id: ")
			scanner.Scan()
			ownerID := scanner.Text()
			var owner *Owner
			for _, o := range owners {
				if o.ID == ownerID {
					owner = o
					break
				}
			}
			if owner == nil {
				log.Println("Enter todo owner name: ")
				scanner.Scan()
				ownerName := scanner.Text()
				owner = &Owner{ID: ownerID, Name: ownerName}
				err := ownerDao.Create(owner)
				if err != nil {
					log.Fatal(err)
				}
			}
			todo.Owner = owner
			todo.OwnerID = owner.ID

			err = todoDao.Create(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo created")
		case "update":
			if todoDao == nil {
				log.Println("Please select a database first")
				break
			}
			log.Print("Enter todo id: ")
			scanner.Scan()
			id := scanner.Text()
			log.Print("Enter todo title: ")
			scanner.Scan()
			title := scanner.Text()
			log.Print("Enter todo completed: ")
			scanner.Scan()
			completed := scanner.Text()
			todo := &Todo{ID: id, Title: title, Completed: completed == "true"}
			err := todoDao.Update(todo)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo updated")
		case "delete":
			if todoDao == nil {
				log.Println("Please select a database first")
				break
			}
			log.Print("Enter todo id: ")
			scanner.Scan()
			id := scanner.Text()
			err := todoDao.Delete(id)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Todo deleted")
		case "done":
			if todoDao == nil {
				log.Println("Please select a database first")
				break
			}
			log.Print("Enter todo id: ")
			scanner.Scan()
			id := scanner.Text()
			err := todoDao.Done(id)
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("Unknown command")
		}
		log.Println("********************************")
		log.Println("Commands:")
		log.Println("pg - select postgres database")
		log.Println("mg - select mongodb database")
		log.Println("list - list all todos")
		log.Println("create - create a todo")
		log.Println("update - update a todo")
		log.Println("delete - delete a todo")
		log.Println("done - mark a todo as done")
		log.Println("quit - quit the program")
		log.Println("********************************")
		log.Print("Enter command: ")
		scanner.Scan()
		cmd = scanner.Text()
	}
	log.Println("Bye!")
}
