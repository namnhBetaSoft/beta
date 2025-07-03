package repository

import (
	"beta/internal/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	User() user.Repository
}

func New(
	db *mongo.Database,
) Registry {
	return &impl{
		user: user.New(db),
	}
}

type impl struct {
	user user.Repository
}

func (r *impl) User() user.Repository {
	return r.user
}
