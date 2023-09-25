package main

import (
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestMGGetAllOwners(t *testing.T) {
	t.Log("Testing Mongo GetAll Owners")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"_id": "test", "name": "test"})
	dao := NewMongoOwnerDao(client)
	owners, err := dao.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(owners) == 0 {
		cleanupMongo(client)
		t.Fatal("expected at least one owner")
	}

	cleanupMongo(client)
}

func TestMGGetOwner(t *testing.T) {
	t.Log("Testing Mongo GetOwner")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"_id": "test", "name": "test"})
	dao := NewMongoOwnerDao(client)
	owner, err := dao.Get("test")
	if err != nil {
		t.Fatal(err)
	}

	if owner == nil {
		cleanupMongo(client)
		t.Fatal("expected an owner")
	}

	cleanupMongo(client)
}

func TestMGCreateOwner(t *testing.T) {
	t.Log("Testing Mongo CreateOwner")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	dao := NewMongoOwnerDao(client)
	owner := &Owner{
		ID:   "test",
		Name: "test",
	}
	err = dao.Create(owner)
	if err != nil {
		t.Fatal(err)
	}

	cleanupMongoOwners(client)
}

func TestMGUpdateOwner(t *testing.T) {
	t.Log("Testing Mongo UpdateOwner")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"_id": "test", "name": "test"})
	dao := NewMongoOwnerDao(client)
	owner := &Owner{
		ID:   "test",
		Name: "testing",
	}
	err = dao.Update(owner)
	if err != nil {
		t.Fatal(err)
	}

	var updatedOwner *Owner
	err = client.Database("todos").Collection("owners").FindOne(ctx, bson.M{"_id": "test"}).Decode(&updatedOwner)
	if err != nil {
		t.Fatal(err)
	}

	if updatedOwner.Name != "testing" {
		t.Fatal("expected owner name to be 'testing'")
	}

	cleanupMongoOwners(client)
}

func TestMGDeleteOwner(t *testing.T) {
	t.Log("Testing Mongo DeleteOwner")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/?directConnection=true"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	client.Database("todos").Collection("owners").InsertOne(ctx, map[string]string{"_id": "test", "name": "test"})
	dao := NewMongoOwnerDao(client)
	err = dao.Delete("test")
	if err != nil {
		t.Fatal(err)
	}

	cleanupMongoOwners(client)
}

func cleanupMongoOwners(client *mongo.Client) {
	ctx := context.Background()
	// delete all todos with id starting with "test"
	client.Database("todos").Collection("owners").DeleteMany(ctx, bson.M{"_id": bson.M{"$regex": "^test"}})
}
