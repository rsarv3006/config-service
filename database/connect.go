package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func Connect() *firestore.Client {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_PROJECT_ID")

	// If GOOGLE_APPLICATION_CREDENTIALS is set, it uses that
	// Otherwise, it uses ADC (gcloud auth or metadata server)
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create Firestore client: %v", err)
	}

	log.Println("Connected to Firestore.")
	return client
}

