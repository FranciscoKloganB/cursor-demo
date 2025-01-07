package utils

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey string

// ActorKey is the string used to set and get the Actor UUID from context.Context.
const ActorKey ctxKey = "actorID"

// GetCtxActor retrieves the Actor UUID from the context or returns uuid.Nil if not set.
func GetCtxActor(ctx context.Context) *uuid.UUID {
	if actorID, ok := ctx.Value(ActorKey).(uuid.UUID); ok {
		return &actorID
	}

	return &uuid.Nil
}

// SetCtxActor sets the Actor UUID in the context. If the UUID is nil, it sets uuid.Nil.
func SetCtxActor(ctx context.Context, actorID *uuid.UUID) context.Context {
	if actorID == nil {
		return context.WithValue(ctx, ActorKey, uuid.Nil)
	}

	return context.WithValue(ctx, ActorKey, *actorID)
}
