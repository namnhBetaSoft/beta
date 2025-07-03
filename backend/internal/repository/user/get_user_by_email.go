package user

import (
	"context"
	"errors"

	"beta/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := i.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, pkgerrors.WithStack(ErrUserNotFound)
		}
		return model.User{}, pkgerrors.WithStack(err)
	}
	return user, nil
}
