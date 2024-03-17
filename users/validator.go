package users

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type UserBodyValidator struct {
	message string
	status  int
}

func (u *UserBodyValidator) Error() string {
	return fmt.Sprintf("%s", u.message)
}

func BodyValidator(u *Users) error {
	ctx := context.TODO()
	if u.Name == "" {
		err := &UserBodyValidator{
			"Name cannot be blank",
			http.StatusBadRequest,
		}
		return err
	}
	if u.Age < 18 {
		err := &UserBodyValidator{
			"User must be at least 18 years old",
			http.StatusBadRequest,
		}
		return err
	}

	existingUser := getCollection().FindOne(ctx, bson.M{"cpf": u.Cpf})

	if existingUser.Err() == nil && existingUser != nil {
		err := &UserBodyValidator{
			"Cpf already registered",
			http.StatusBadRequest,
		}

		return err
	}
	return nil
}
