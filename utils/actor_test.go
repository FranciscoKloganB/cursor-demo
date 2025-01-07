//go:build unit

package utils_test

import (
	"context"
	"testing"

	"encore.app/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestActorKey(t *testing.T) {
	assert.Equal(t, string(utils.ActorKey), "actorID")
}

func TestGetCtxActor(t *testing.T) {
	t.Run("returns uuid.Nil when actor is not set", func(t *testing.T) {
		ctx := context.Background()
		actor := utils.GetCtxActor(ctx)
		assert.Equal(t, uuid.Nil, *actor)
	})

	t.Run("returns actor UUID when set", func(t *testing.T) {
		expectedActor := uuid.New()
		ctx := context.WithValue(context.Background(), utils.ActorKey, expectedActor)
		actor := utils.GetCtxActor(ctx)
		assert.Equal(t, expectedActor, *actor)
	})
}

func TestSetCtxActor(t *testing.T) {
	t.Run("sets uuid.Nil when actor is nil", func(t *testing.T) {
		ctx := context.Background()
		ctx = utils.SetCtxActor(ctx, nil)
		actor, ok := ctx.Value(utils.ActorKey).(uuid.UUID)
		assert.True(t, ok)
		assert.Equal(t, uuid.Nil, actor)
	})

	t.Run("sets actor UUID when provided", func(t *testing.T) {
		expectedActor := uuid.New()
		ctx := context.Background()
		ctx = utils.SetCtxActor(ctx, &expectedActor)
		actor, ok := ctx.Value(utils.ActorKey).(uuid.UUID)
		assert.True(t, ok)
		assert.Equal(t, expectedActor, actor)
	})
}
