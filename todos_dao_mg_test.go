package main

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestMGGetAll(t *testing.T) {
	t.Log("Testing Mongo GetAll")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})
	client.Database("todos").Collection("todos").InsertOne(ctx, map[string]interface{}{
		"id":         "test",
		"title":      "test",
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   "test",
	})

	dao := NewMongoTodoDao(client)
	todos, err := dao.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) == 0 {
		cleanupMongo(client)
		t.Fatal("expected at least one todo")
	}

	cleanupMongo(client)
}

func TestMGGet(t *testing.T) {
	t.Log("Testing Mongo Get")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})
	client.Database("todos").Collection("todos").InsertOne(ctx, map[string]interface{}{
		"id":         "test",
		"title":      "test",
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   "test",
	})

	dao := NewMongoTodoDao(client)
	todo, err := dao.Get("test")
	if err != nil {
		cleanupMongo(client)
		t.Fatal(err)
	}

	if todo == nil {
		cleanupMongo(client)
		t.Fatal("expected at least one todo")
	}

	if todo.Title != "test" {
		cleanupMongo(client)
		t.Fatal("expected todo title to be 'test'")
	}

	cleanupMongo(client)
}

func TestMGCreate(t *testing.T) {
	t.Log("Testing Mongo Create")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})

	dao := NewMongoTodoDao(client)
	todo := &Todo{ID: "test", Title: "test", Owner: &Owner{ID: "test"}}
	err = dao.Create(todo)
	if err != nil {
		cleanupMongo(client)
		t.Fatal(err)
	}

	cleanupMongo(client)
}

func TestMGUpdate(t *testing.T) {
	t.Log("Testing Mongo Update")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})
	client.Database("todos").Collection("todos").InsertOne(ctx, map[string]interface{}{
		"id":         "test",
		"title":      "test",
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   "test",
	})

	dao := NewMongoTodoDao(client)
	todo := &Todo{ID: "test", Title: "test", Completed: true}
	err = dao.Update(todo)
	if err != nil {
		cleanupMongo(client)
		t.Fatal(err)
	}

	cleanupMongo(client)
}

func TestMGDelete(t *testing.T) {
	t.Log("Testing Mongo Delete")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})
	client.Database("todos").Collection("todos").InsertOne(ctx, map[string]interface{}{
		"id":         "test",
		"title":      "test",
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   "test",
	})

	dao := NewMongoTodoDao(client)
	err = dao.Delete("test")
	if err != nil {
		cleanupMongo(client)
		t.Fatal(err)
	}

	cleanupMongo(client)
}

func TestMGDone(t *testing.T) {
	t.Log("Testing Mongo Done")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"id": "test", "name": "test"})
	client.Database("todos").Collection("todos").InsertOne(ctx, map[string]interface{}{
		"id":         "test",
		"title":      "test",
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   "test",
	})

	dao := NewMongoTodoDao(client)
	err = dao.Done("test")
	if err != nil {
		cleanupMongo(client)
		t.Fatal(err)
	}

	todo := &Todo{}
	client.Database("todos").Collection("todos").FindOne(ctx, map[string]string{"id": "test"}).Decode(todo)
	if !todo.Completed {
		cleanupMongo(client)
		t.Fatal("expected todo to be completed")
	}

	cleanupMongo(client)
}

func cleanupMongo(client *mongo.Client) {
	ctx := context.Background()
	client.Database("todos").Collection("todos").DeleteMany(ctx, map[string]string{"id": "test"})
	client.Database("todos").Collection("owners").DeleteMany(ctx, map[string]string{"id": "test"})
}
