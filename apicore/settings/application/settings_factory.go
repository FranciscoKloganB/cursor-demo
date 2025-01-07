package application

import (
	"fmt"
	"time"

	"encore.app/apicore/settings/application/commands"
	"encore.app/apicore/settings/domain/entities"
	"github.com/google/uuid"
)

// CreateSetting creates a new Setting from the given struct
//
// When the input struct can not be converted to a Setting domain objec it will return an errort.
func createSettingFactory(value interface{}) (entities.Setting, error) {
	switch v := value.(type) {
	case commands.CreateSettingCommand:
		actor := *v.ActorID
		now := time.Now().UTC()
		uid := uuid.New()

		setting, err := entities.NewSetting(
			uid,
			v.Name,
			v.Slug,
			v.Hint,
			v.IsActive,
			now,
			actor,
			nil,
			&uuid.Nil,
			nil,
			&uuid.Nil,
			1,
		)

		if err != nil {
			return setting, err
		}

		return setting, nil
	default:
		return entities.Setting{}, fmt.Errorf("unknown value type: %T can not be converted to domain entity of type Setting", v)
	}
}
