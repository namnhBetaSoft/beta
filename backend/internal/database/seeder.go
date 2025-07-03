package database

import (
	"beta/internal/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Seeder struct {
	db *mongo.Database
}

func NewSeeder(db *mongo.Database) *Seeder {
	return &Seeder{db: db}
}

// SeedUsers adds sample users to the database
func (s *Seeder) SeedUsers(ctx context.Context) error {
	collection := s.db.Collection("users")

	// Check if users already exist
	count, err := collection.CountDocuments(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Users already exist in database, skipping seeding")
		return nil
	}

	// Hash passwords
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	// Sample users
	users := []model.User{
		{
			ID:        primitive.NewObjectID(),
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  string(hashedPassword1),
			Image:     "https://example.com/avatar1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Password:  string(hashedPassword2),
			Image:     "https://example.com/avatar2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Admin User",
			Email:     "admin@example.com",
			Password:  string(hashedPassword2),
			Image:     "https://example.com/admin.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Insert users
	for _, user := range users {
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			log.Printf("Error inserting user %s: %v", user.Email, err)
			return err
		}
		log.Printf("Successfully inserted user: %s", user.Email)
	}

	log.Printf("Successfully seeded %d users", len(users))
	return nil
}

// SeedAll runs all seeding operations
func (s *Seeder) SeedAll(ctx context.Context) error {
	log.Println("Starting database seeding...")

	if err := s.SeedUsers(ctx); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully!")
	return nil
}
