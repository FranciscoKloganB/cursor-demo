package application

import (
	"fmt"
	"time"

	"encore.app/apicore/organisations/application/commands"
	"encore.app/apicore/organisations/domain/entities"
	"github.com/google/uuid"
)

// createOrganisationFactory creates a new Organisation from the given struct
//
// When the input struct cannot be converted to an Organisation domain object, it will return an error.
func createOrganisationFactory(value interface{}) (entities.Organisation, error) {
	switch v := value.(type) {
	case commands.CreateOrganisationCommand:
		actor := *v.ActorID
		now := time.Now().UTC()
		uid := uuid.New()

		organisation, err := entities.NewOrganisation(
			uid,
			v.Name,
			now,
			actor,
			nil,
			&uuid.Nil,
			nil,
			&uuid.Nil,
			1,
		)

		if err != nil {
			return organisation, err
		}

		return organisation, nil
	default:
		return entities.Organisation{}, fmt.Errorf("unknown value type: %T cannot be converted to domain entity of type Organisation", v)
	}
}
