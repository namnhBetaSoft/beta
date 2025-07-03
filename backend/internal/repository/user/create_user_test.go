package user_test

import (
	"beta/internal/model"
	"beta/internal/repository/user"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	// Initialize the mock MongoDB test environment
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// Define test cases
	tests := []struct {
		name          string
		user          model.User
		mockResponses func(mt *mtest.T)
		expectError   bool
		errorContains string
	}{
		{
			name: "successful user creation",
			user: model.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "hashedpassword123",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne success
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			expectError: false,
		},
		{
			name: "successful user creation with all fields",
			user: model.User{
				Name:     "Jane Smith",
				Email:    "jane@example.com",
				Password: "password456",
				Image:    "https://example.com/avatar.jpg",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne success
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			expectError: false,
		},
		{
			name: "database insert failure - duplicate key",
			user: model.User{
				Name:     "Duplicate User",
				Email:    "duplicate@example.com",
				Password: "password789",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne with duplicate key error
				mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
					Index:   0,
					Code:    11000,
					Message: "duplicate key error",
				}))
			},
			expectError:   true,
			errorContains: "duplicate key error",
		},
		{
			name: "database insert failure - write exception",
			user: model.User{
				Name:     "Failed User",
				Email:    "failed@example.com",
				Password: "password123",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne with write exception
				mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Code:    1,
					Message: "write exception",
				}))
			},
			expectError:   true,
			errorContains: "write exception",
		},
		{
			name: "user with empty name",
			user: model.User{
				Name:     "",
				Email:    "empty@example.com",
				Password: "password123",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne success
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			expectError: false,
		},
		{
			name: "user with empty email",
			user: model.User{
				Name:     "No Email User",
				Email:    "",
				Password: "password123",
			},
			mockResponses: func(mt *mtest.T) {
				// Mock InsertOne success
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			// Setup the repository using the actual New() function with mock database
			repo := user.New(mt.DB)

			// Keep original user state for comparison
			originalUser := tt.user
			originalID := tt.user.ID
			originalCreatedAt := tt.user.CreatedAt
			originalUpdatedAt := tt.user.UpdatedAt

			// Mock the database responses
			tt.mockResponses(mt)

			// Call the CreateUser method
			err := repo.CreateUser(context.Background(), tt.user)

			// Validate the error
			if tt.expectError {
				assert.Error(t, err, "Expected error but got none for test: %s", tt.name)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains, "Error message should contain expected text for test: %s", tt.name)
				}
			} else {
				assert.NoError(t, err, "Expected no error but got: %v for test: %s", err, tt.name)
			}

			// Validate that the original user object is not modified
			// Since Go passes structs by value, the original should remain unchanged
			assert.Equal(t, originalUser.Name, tt.user.Name, "User name should not be modified")
			assert.Equal(t, originalUser.Email, tt.user.Email, "User email should not be modified")
			assert.Equal(t, originalUser.Password, tt.user.Password, "User password should not be modified")
			assert.Equal(t, originalUser.Image, tt.user.Image, "User image should not be modified")
			assert.Equal(t, originalID, tt.user.ID, "User ID should not be modified in original object")
			assert.Equal(t, originalCreatedAt, tt.user.CreatedAt, "CreatedAt should not be modified in original object")
			assert.Equal(t, originalUpdatedAt, tt.user.UpdatedAt, "UpdatedAt should not be modified in original object")
		})
	}
}

// TestCreateUserContextCancellation tests context cancellation behavior
func TestCreateUserContextCancellation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	tests := []struct {
		name          string
		user          model.User
		contextSetup  func() (context.Context, context.CancelFunc)
		expectedError bool
	}{
		{
			name: "context canceled before operation",
			user: model.User{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			contextSetup: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel() // Cancel immediately
				return ctx, cancel
			},
			expectedError: true,
		},
		{
			name: "context with timeout",
			user: model.User{
				Name:     "Timeout User",
				Email:    "timeout@example.com",
				Password: "password123",
			},
			contextSetup: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 0) // Immediate timeout
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			// Setup the repository
			repo := user.New(mt.DB)

			// Setup context
			ctx, cancel := tt.contextSetup()
			defer cancel()

			// Call the CreateUser method
			err := repo.CreateUser(ctx, tt.user)

			// Validate the error
			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "context")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
