package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetDB() *mongo.Database
}

type service struct {
	db *mongo.Client
}

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_ROOT_PASSWORD")
)

func New() Service {
	var uri string

	// Build MongoDB URI with or without authentication
	if username != "" && password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", username, password, host, port, database)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	}

	log.Printf("Connecting to MongoDB: %s", uri)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetDB() *mongo.Database {
	return s.db.Database(database)
}
