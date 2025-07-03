package user

import (
	"beta/internal/model"
	"context"
	"time"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUserFromOAuth creates a user from OAuth profile.
func (i impl) CreateUserFromOAuth(ctx context.Context, profile OAuthProfile) (model.User, error) {
	now := time.Now()
	user := model.User{
		ID:            primitive.NewObjectID(),
		Name:          profile.Name,
		Email:         profile.Email,
		Image:         profile.Picture,
		EmailVerified: now, // OAuth users are considered verified
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := i.CreateUser(ctx, user); err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	return user, nil
}

// FindByOAuthProvider retrieves a user by their OAuth provider and account ID.
func (i impl) FindByOAuthProvider(ctx context.Context, provider, providerAccountID string) (model.User, error) {
	var user model.User
	filter := bson.M{
		"oauthProvider":  provider,
		"oauthAccountId": providerAccountID,
	}

	err := i.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if pkgerrors.Is(err, mongo.ErrNoDocuments) {
			return user, ErrUserNotFound
		}
		return user, pkgerrors.WithStack(err)
	}
	return user, nil
}
