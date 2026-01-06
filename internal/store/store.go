package store

import (
	"context"
	"github.com/google/uuid"
)

type WebhookEvent struct {
	ID		uuid.UUID
	Source	string
	Payload	[]byte
	Status	string
	Attempt	int
}

type WebhookStore interface {
	Insert(ctx context.Context, e WebhookEvent) error
}