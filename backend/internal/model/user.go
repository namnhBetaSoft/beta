package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	EmailVerified time.Time          `json:"emailVerified,omitempty" bson:"emailVerified,omitempty"`
	Image         string             `json:"image,omitempty" bson:"image,omitempty"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Account represents an OAuth account linked to a user
type Account struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID            primitive.ObjectID `json:"userId" bson:"userId"`
	Type              string             `json:"type" bson:"type"`
	Provider          string             `json:"provider" bson:"provider"`
	ProviderAccountID string             `json:"providerAccountId" bson:"providerAccountId"`
	RefreshToken      string             `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	AccessToken       string             `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	ExpiresAt         int                `json:"expiresAt,omitempty" bson:"expiresAt,omitempty"`
	TokenType         string             `json:"tokenType,omitempty" bson:"tokenType,omitempty"`
	Scope             string             `json:"scope,omitempty" bson:"scope,omitempty"`
	IDToken           string             `json:"idToken,omitempty" bson:"idToken,omitempty"`
	SessionState      string             `json:"sessionState,omitempty" bson:"sessionState,omitempty"`
	CreatedAt         time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Session represents a user session
type Session struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SessionToken string             `json:"sessionToken" bson:"sessionToken"`
	UserID       primitive.ObjectID `json:"userId" bson:"userId"`
	Expires      time.Time          `json:"expires" bson:"expires"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// VerificationToken represents a token used for email verification
type VerificationToken struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Identifier string             `json:"identifier" bson:"identifier"`
	Token      string             `json:"token" bson:"token"`
	Expires    time.Time          `json:"expires" bson:"expires"`
}
