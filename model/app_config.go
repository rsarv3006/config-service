package model

import "time"

type AppConfig struct {
	ID        string         `firestore:"id" json:"id"`
	CreatedAt time.Time      `firestore:"created_at" json:"created_at"`
	AppName   string         `firestore:"app_name" json:"app_name"`
	Version   int            `firestore:"version" json:"version"`
	Status    string         `firestore:"status" json:"status"`
	Config    map[string]any `firestore:"config" json:"config"`
}
