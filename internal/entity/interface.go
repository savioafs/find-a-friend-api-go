package entity

type PetStorer interface {
	Save(pet *Pet) error
	FindPetByID(id string) (*Pet, error)
	FindByCity(city string) ([]Pet, error)
	AllByOrganization(org_id string) ([]Pet, error)
	AllByCharacteristics(characteristics []string) ([]Pet, error)
	Update(id string, pet *Pet) (*Pet, error)
	Delete(id string) error
}

type OrganizationStorer interface {
	Save(org *Organization) error
	FindOrganizationByID(id string) (*Organization, error)
	FindOrganizationByEmail(email string) (*Organization, error)
	UpdateOrganization(id string) (*Organization, error)
	DeleteOrganization(id string) error
}
