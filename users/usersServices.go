package users

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUsers() ([]Users, error) {
	ctx := context.TODO()
	var us []Users
	u, err := FindUsers()
	if err != nil {
		fmt.Println(err)
	}

	if u.Next(ctx) {
		if err = u.All(ctx, &us); err != nil {
			return nil, fmt.Errorf("error decoding users")
		}
	}

	return us, nil
}

func CreateUsers(u Users) *Users {
	user := InsertUser(u)

	return user
}

func UpdateUsersName(id string, nn string) (*Users, error) {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("hex string not valid")
	}
	m := mongo.NewUpdateOneModel()

	f := bson.D{{"_id", i}}
	u := bson.D{{"$set", bson.D{{"name", nn}}}}

	m.SetFilter(f)
	m.SetUpdate(u)

	dc, err := Update(*m)
	if err != nil {
		return nil, fmt.Errorf("error on updating user")
	}

	return dc, nil
}
