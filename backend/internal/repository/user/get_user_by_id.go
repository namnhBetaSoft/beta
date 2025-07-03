package user

import (
	"beta/internal/model"
	"context"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByID retrieves a user by their unique identifier (ID).
func (i impl) GetByID(ctx context.Context, userID primitive.ObjectID) (model.User, error) {
	var user model.User
	filter := bson.M{"_id": userID}

	err := i.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if pkgerrors.Is(err, mongo.ErrNoDocuments) {
			return user, ErrUserNotFound
		}
		return user, pkgerrors.WithStack(err)
	}
	return user, nil
}
