package user

import (
	"beta/internal/model"
	"context"
	"time"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is a method to create a new user in the database.
func (i impl) CreateUser(ctx context.Context, user model.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := i.collection.InsertOne(ctx, user)
	return pkgerrors.WithStack(err)
}
