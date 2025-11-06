package model

import "time"

type User struct {
	ID        string    `firestore:"id" json:"id"`
	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
	Role      string    `firestore:"role" json:"role"`
	AppName   string    `firestore:"app_name" json:"app_name"`
	Status    string    `firestore:"status" json:"status"`
}
