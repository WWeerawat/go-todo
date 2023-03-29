package todo

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go-todo/config"
	"go-todo/models"
	"go-todo/todo/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoCollection *mongo.Collection = config.GetCollection(config.DB, "todo")
var validate = validator.New()

func Create(ctx context.Context, newTodo models.ToDo) (*mongo.InsertOneResult, error) {
	result, err := todoCollection.InsertOne(ctx, newTodo)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func GetAll(ctx context.Context) ([]models.ToDo, error) {
	var todos []models.ToDo
	cursor, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {

		return nil, err
	}

	if err = cursor.All(ctx, &todos); err != nil {

		return nil, err
	}

	return todos, nil
}

func GetById(ctx context.Context, id string) (*models.ToDo, error) {
	var todo *models.ToDo
	objId, _ := primitive.ObjectIDFromHex(id)

	err := todoCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func GetByUser(ctx context.Context, userId string) (*models.ToDo, error) {
	var todo *models.ToDo

	err := todoCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func Update(ctx context.Context, id string, input dto.UpdateToDoInput) (*mongo.UpdateResult, error) {

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := validate.Struct(&input); err != nil {
		return nil, err
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", input}}

	result, err := todoCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Delete(ctx context.Context, id string) (*models.ToDo, error) {
	var todo *models.ToDo
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := todoCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	if result.DeletedCount < 1 {
		return nil, err
	}

	return todo, nil
}
