package user

import (
	"beta/internal/model"
	"context"
	"time"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser creates a new user in the database.
func (i impl) CreateUser(ctx context.Context, user model.User) error {
	// Set timestamps if not provided
	now := time.Now()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	// Generate ID if not provided
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := i.collection.InsertOne(ctx, user)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
