package entity

type PetStorer interface {
	Save(pet *Pet) error
	FindPetByID(id string) (*Pet, error)
	// FindByCity(city string) ([]entity.Pet, error)
	// AllByOrganization(org_id string) ([]entity.Pet, error)
	// AllByCharacteristics(characteristics []string) ([]entity.Pet, error)
	// Update(id string) (*entity.Pet, error)
	// Delete(id string) error
}

type OrganizationStorer interface {
	// Save(org *entity.Organization) error
	FindOrganizationByID(id string) (*Organization, error)
	// UpdateOrganization(id string) (*entity.Organization, error)
	// DeleteOrganization(id string) error
}