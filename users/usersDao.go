package users

import (
	"basic-mongo-crud/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
	ops options.FindOneAndUpdateOptions
)

func getCollection() *mongo.Collection {
	return db.GetCollection("users")
}

func FindUsers() (*mongo.Cursor, error) {
	c, err := getCollection().Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	if !c.Next(ctx) {
		return nil, fmt.Errorf("empty collection")
	}

	return c, nil
}

func InsertUser(u Users) *Users {
	_, err := getCollection().InsertOne(ctx, u)

	if err != nil {
		fmt.Println(err)
	}

	return &u
}

func DeleteUser(u UserCpf) (string, error) {
	if r := getCollection().FindOneAndDelete(ctx, bson.D{{"cpf", u.Cpf}}); r.Err() != nil {
		return "", fmt.Errorf("error deleting user")
	}

	return "user deleted successfully", nil
}

func Update(update mongo.UpdateOneModel) (*Users, error) {
	var ud *Users
	ops.SetReturnDocument(options.After)
	err := getCollection().FindOneAndUpdate(ctx, update.Filter, update.Update, &ops).Decode(&ud)
	if err != nil {
		return nil, fmt.Errorf("error updating user")
	}

	return ud, nil
}
