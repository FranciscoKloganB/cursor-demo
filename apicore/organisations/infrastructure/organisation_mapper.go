package infrastructure

import (
	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/organisations/domain/entities"
)

// ToPersistence converts a domain organisation to a persistence entity.
func toPersistence(organisation entities.Organisation) (db.Organisation, error) {
	return db.Organisation{
		ID:        organisation.ID,
		Name:      organisation.Name,
		CreatedAt: organisation.CreatedAt,
		CreatedBy: organisation.CreatedBy,
		DeletedAt: organisation.DeletedAt,
		DeletedBy: organisation.DeletedBy,
		UpdatedAt: organisation.UpdatedAt,
		UpdatedBy: organisation.UpdatedBy,
		Version:   organisation.Version,
	}, nil
}

// ToDomain converts a persistence entity to a domain organisation.
func toDomain(entity db.Organisation) (entities.Organisation, error) {
	organisation, err := entities.NewOrganisation(
		entity.ID,
		entity.Name,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.DeletedAt,
		entity.DeletedBy,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.Version,
	)

	return organisation, err
}
