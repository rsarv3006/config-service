package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
	AppName   string    `json:"app_name"`
}
