package user

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go-todo/config"
	"go-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func Create(ctx context.Context, newUser models.User) (*mongo.InsertOneResult, error) {
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.D{})
	if err != nil {

		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {

		return nil, err
	}

	return users, nil
}

func GetById(ctx context.Context, id string) (*models.User, error) {
	var user *models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Update(ctx context.Context, id string, update bson.M) (*models.User, error) {
	var user *models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	if err := validate.Struct(&user); err != nil {
		return nil, err
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func Delete(ctx context.Context, id string) (*models.User, error) {
	var user *models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	if result.DeletedCount < 1 {
		return nil, err
	}

	return user, nil
}
