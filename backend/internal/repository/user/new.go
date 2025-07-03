package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"beta/internal/model"
)

const (
	CollectionName = "users"
)

// OAuthProfile represents user profile from OAuth provider
type OAuthProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	Provider string `json:"provider"`
}

// Repository interface defines all user repository methods
type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetByID(ctx context.Context, userID primitive.ObjectID) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	CreateUserFromOAuth(ctx context.Context, profile OAuthProfile) (model.User, error)
	UpdateEmailVerified(ctx context.Context, email string) error
	UpdateUserInfo(ctx context.Context, email string, name string, password string) error
	FindByOAuthProvider(ctx context.Context, provider, providerAccountID string) (model.User, error)
}

// impl is the implementation of Repository interface
type impl struct {
	collection *mongo.Collection
}

func New(db *mongo.Database) Repository {
	collection := db.Collection(CollectionName)

	// Create indexes for better performance
	ctx := context.Background()

	// Unique index on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Compound index for OAuth provider search
	oauthIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "oauthProvider", Value: 1},
			{Key: "oauthAccountId", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}

	// Create indexes (ignore errors if they already exist)
	collection.Indexes().CreateMany(ctx, []mongo.IndexModel{emailIndex, oauthIndex})

	return &impl{
		collection: collection,
	}
}
