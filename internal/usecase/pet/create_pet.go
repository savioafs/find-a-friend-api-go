package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type CreatePetUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewCreatePetUseCase(repo entity.PetStorer, repoOrg entity.OrganizationStorer) *CreatePetUseCase {
	return &CreatePetUseCase{
		PetRepository: repo,
		OrgRepository: repoOrg,
	}
}

func (u *CreatePetUseCase) Execute(input dto.PetInputDTO) (dto.PetOutputDTO, error) {
	pet, err := entity.NewPet(
		input.OrgID,
		input.Name,
		input.About,
		input.Age,
		input.Size,
		input.EnergyLevel,
		input.DependencyLevel,
		input.Environment,
		input.Photos,
		input.RequirementsForAdoption)

	if err != nil {
		return dto.PetOutputDTO{}, err
	}

	err = u.PetRepository.Save(pet)
	if err != nil {
		return dto.PetOutputDTO{}, err
	}

	org, err := u.OrgRepository.FindOrganizationByID(pet.OrgID)
	if err != nil {
		return dto.PetOutputDTO{}, err
	}

	petOutput := dto.PetOutputDTO{
		ID:                      pet.ID,
		Org:                     pet.OrgID,
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
