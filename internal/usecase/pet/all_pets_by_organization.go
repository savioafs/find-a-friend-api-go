package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type AllPetsByOrganizationUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewAllPetsByOrganizationUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *AllPetsByOrganizationUseCase {
	return &AllPetsByOrganizationUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *AllPetsByOrganizationUseCase) Execute(organizationID string) (dto.PetOutputDTO, error) {
	return dto.PetOutputDTO{}, nil
}
