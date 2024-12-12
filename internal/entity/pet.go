package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrOrgIDIsRequired                   = errors.New("org id is required")
	ErrNameIsRequired                    = errors.New("name is required")
	ErrAboutIsRequired                   = errors.New("about is required")
	ErrAgeIsRequired                     = errors.New("age is required")
	ErrSizeIsRequired                    = errors.New("size is required")
	ErrEnergyLevelIsRequired             = errors.New("energy level is required")
	ErrDependencyLevelIsRequired         = errors.New("dependency level is required")
	ErrEnvironmentIsRequired             = errors.New("enviroment is required")
	ErrPhotosIsRequired                  = errors.New("photos are required")
	ErrRequirementsForAdoptionIsRequired = errors.New("requirements are required")
)

type Pet struct {
	ID                      string   `db:"id"`
	OrgID                   string   `db:"org_id"`
	Name                    string   `db:"name"`
	About                   string   `db:"about"`
	Age                     string   `db:"age"`
	Size                    string   `db:"size"`
	EnergyLevel             string   `db:"energy_level"`
	DependencyLevel         string   `db:"dependency_level"`
	Environment             string   `db:"enviroment"`
	Photos                  []string `db:"photos"`
	RequirementsForAdoption []string `db:"requirements_for_adoption"`
}

func NewPet(orgID, name, about, age, size, energyLevel, dependencyLevel, enviroment string, photos, requirements []string) (*Pet, error) {
	pet := &Pet{
		ID:                      uuid.New().String(),
		OrgID:                   orgID,
		Name:                    name,
		About:                   about,
		Age:                     age,
		Size:                    size,
		EnergyLevel:             energyLevel,
		DependencyLevel:         dependencyLevel,
		Environment:             enviroment,
		Photos:                  photos,
		RequirementsForAdoption: requirements,
	}

	err := pet.Validate()
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (p *Pet) Validate() error {
	if p.OrgID == "" {
		return ErrOrgIDIsRequired
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.About == "" {
		return ErrAboutIsRequired
	}

	if p.Age == "" {
		return ErrAgeIsRequired
	}

	if p.Size == "" {
		return ErrSizeIsRequired
	}

	if p.EnergyLevel == "" {
		return ErrEnergyLevelIsRequired
	}

	if p.DependencyLevel == "" {
		return ErrDependencyLevelIsRequired
	}

	if p.Environment == "" {
		return ErrEnvironmentIsRequired
	}

	if p.Photos == nil {
		return ErrPhotosIsRequired
	}

	if p.RequirementsForAdoption == nil {
		return ErrRequirementsForAdoptionIsRequired
	}

	return nil
}
