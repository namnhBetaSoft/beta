package main

import (
	"beta/internal/database"
	"context"
	"log"
	"time"
)

func main() {
	log.Println("Starting database seeding...")

	// Initialize database service
	dbService := database.New()

	// Get database connection
	db := dbService.GetDB()

	// Create seeder
	seeder := database.NewSeeder(db)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Run seeding
	if err := seeder.SeedAll(ctx); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("Database seeding completed successfully!")
}
