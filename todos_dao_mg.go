package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoDaoMongoImpl struct {
	client *mongo.Client
}

func NewMongoTodoDao(client *mongo.Client) TodoDao {
	return &TodoDaoMongoImpl{client}
}

func (dao *TodoDaoMongoImpl) GetAll() ([]*Todo, error) {
	ctx := context.Background()
	todos := []*Todo{}
	collection := dao.client.Database("todos").Collection("todos")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (dao *TodoDaoMongoImpl) Get(id string) (*Todo, error) {
	ctx := context.Background()
	todo := &Todo{}
	collection := dao.client.Database("todos").Collection("todos")
	if err := collection.FindOne(ctx, bson.M{"id": id}).Decode(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *TodoDaoMongoImpl) Create(todo *Todo) error {

	return errors.New("not implemented")
}

func (dao *TodoDaoMongoImpl) Update(todo *Todo) error {
	return errors.New("not implemented")
}

func (dao *TodoDaoMongoImpl) Delete(id string) error {
	return errors.New("not implemented")
}

func (dao *TodoDaoMongoImpl) Done(id string) error {
	return errors.New("not implemented")
}

func (dao *TodoDaoMongoImpl) GetOwners() ([]*Owner, error) {
	return nil, errors.New("not implemented")
}
