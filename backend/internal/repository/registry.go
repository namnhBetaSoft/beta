package repository

import "go.mongodb.org/mongo-driver/mongo"

type Registry interface {
	// Add methods that will be implemented by the repository
}

func New(
	db *mongo.Database,
) Registry {
	return &impl{}
}

type impl struct {
	// Add fields that will be used in the repository implementation
}
