// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type BotConfig struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Securities json.RawMessage
	ApiKey     string
}
