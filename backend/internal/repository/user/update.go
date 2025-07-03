package user

import (
	"context"
	"time"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateEmailVerified updates the email verification status of a user.
func (i impl) UpdateEmailVerified(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"emailVerified": time.Now(), "updatedAt": time.Now()}}

	result, err := i.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

// UpdateUserInfo updates the user's information such as name and password.
func (i impl) UpdateUserInfo(ctx context.Context, email string, name string, password string) error {
	filter := bson.M{"email": email}
	updateFields := bson.M{
		"name":      name,
		"updatedAt": time.Now(),
	}

	// Only update password if provided
	if password != "" {
		updateFields["password"] = password
	}

	update := bson.M{"$set": updateFields}
	result, err := i.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
