package user

import (
	"beta/internal/model"
	"context"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUserByEmail retrieves a user by their email address.
func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	filter := bson.M{"email": email}

	err := i.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if pkgerrors.Is(err, mongo.ErrNoDocuments) {
			return user, ErrUserNotFound
		}
		return user, pkgerrors.WithStack(err)
	}
	return user, nil
}
