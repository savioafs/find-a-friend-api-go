package usecase

import (
	"errors"

	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type FindPetByIDUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewFindPetByIDUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *FindPetByIDUseCase {
	return &FindPetByIDUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *FindPetByIDUseCase) Execute(id string) (dto.PetOutputDTO, error) {
	if id == "" {
		return dto.PetOutputDTO{}, errors.New("id is required for search")
	}

	pet, err := u.PetRepository.FindPetByID(id)
	if err != nil {
		return dto.PetOutputDTO{}, err
	}

	org, err := u.OrgRepository.FindOrganizationByID(pet.OrgID)
	if err != nil {
		return dto.PetOutputDTO{}, err
	}

	petOutput := dto.PetOutputDTO{
		ID:                      pet.ID,
		Org:                     org.Name,
		OrgWhatspp:              org.Whatsapp,
		Name:                    pet.Name,
		About:                   pet.About,
		Age:                     pet.Age,
		Size:                    pet.Size,
		EnergyLevel:             pet.EnergyLevel,
		DependencyLevel:         pet.DependencyLevel,
		Environment:             pet.Environment,
		Photos:                  pet.Photos,
		RequirementsForAdoption: pet.RequirementsForAdoption,
	}

	return petOutput, nil
}
