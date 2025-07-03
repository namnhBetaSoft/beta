package user

import (
	"beta/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionName = "users"
)

// OAuthProfile represents user profile from OAuth provider
type OAuthProfile struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetByID(ctx context.Context, userID primitive.ObjectID) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	CreateUserFromOAuth(ctx context.Context, profile OAuthProfile) (model.User, error)
	UpdateEmailVerified(ctx context.Context, email string) error
	UpdateUserInfo(ctx context.Context, email string, name string, password string) error
	FindByOAuthProvider(ctx context.Context, provider, providerAccountID string) (model.User, error)
}

type impl struct {
	collection *mongo.Collection
}

func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i impl) GetByID(ctx context.Context, userID primitive.ObjectID) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i impl) CreateUserFromOAuth(ctx context.Context, profile OAuthProfile) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i impl) UpdateEmailVerified(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (i impl) UpdateUserInfo(ctx context.Context, email string, name string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (i impl) FindByOAuthProvider(ctx context.Context, provider, providerAccountID string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func New(db *mongo.Database) Repository {
	return &impl{
		collection: db.Collection(CollectionName),
	}
}
