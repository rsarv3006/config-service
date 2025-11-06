package database

import (
	"config-service/auth"
	"config-service/model"
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

const usersCollection = "ConfigService-users"

func CreateUserAccounts(client *firestore.Client, jwtSecret string) {
	log.Println("Creating admin account")
	adminAccountToken, err := createAdminAccount(client, jwtSecret)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(adminAccountToken)
	log.Println("Admin account created")

	appAccounts, err := CreateAppAccounts(client, jwtSecret)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(appAccounts)
	log.Println("App accounts created")
}

func createAdminAccount(client *firestore.Client, jwtSecret string) (string, error) {
	ctx := context.Background()

	// Check if admin user already exists
	iter := client.Collection(usersCollection).
		Where("role", "==", "admin").
		Limit(1).
		Documents(ctx)
	defer iter.Stop()

	_, err := iter.Next()
	if err != iterator.Done {
		// Admin already exists or there was an error
		if err != nil {
			return "", err
		}
		return "", nil
	}

	// Create admin account
	adminAccount := model.User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Role:      "admin",
		AppName:   "",
		Status:    "active",
	}

	_, err = client.Collection(usersCollection).Doc(adminAccount.ID).Set(ctx, adminAccount)
	if err != nil {
		return "", err
	}

	return auth.GenerateJWTFromSecret(&adminAccount, jwtSecret)
}

var appAccounts = []string{"basketbuddy", "wishlistwrangler", "bedtimebully", "daysuntilwhen", "chorddetector"}

type CreatedAppAccount struct {
	Name  string
	Token string
}

func CreateAppAccounts(client *firestore.Client, jwtSecret string) ([]CreatedAppAccount, error) {
	log.Println("Creating app accounts")
	ctx := context.Background()
	var errorFromProcess error
	createdAccounts := []CreatedAppAccount{}

	for _, appName := range appAccounts {
		// Check if app user already exists
		iter := client.Collection(usersCollection).
			Where("app_name", "==", appName).
			Limit(1).
			Documents(ctx)

		_, err := iter.Next()
		iter.Stop()

		if err != iterator.Done {
			// User already exists or there was an error
			if err != nil {
				errorFromProcess = err
			}
			continue
		}

		// Create app account
		appAccount := model.User{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			Role:      "app",
			AppName:   appName,
			Status:    "active",
		}

		_, err = client.Collection(usersCollection).Doc(appAccount.ID).Set(ctx, appAccount)
		if err != nil {
			errorFromProcess = err
			continue
		}

		token, err := auth.GenerateJWTFromSecret(&appAccount, jwtSecret)
		if err != nil {
			errorFromProcess = err
			continue
		}

		createdAccounts = append(createdAccounts, CreatedAppAccount{
			Name:  appName,
			Token: token,
		})
	}

	return createdAccounts, errorFromProcess
}
