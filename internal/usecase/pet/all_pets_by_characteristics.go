package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type AllPetsByCharacteristicsUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewAllPetsByCharacteristicsUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *AllPetsByCharacteristicsUseCase {
	return &AllPetsByCharacteristicsUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *AllPetsByCharacteristicsUseCase) Execute(characteristics []string) (dto.PetOutputDTO, error) {
	return dto.PetOutputDTO{}, nil
}
