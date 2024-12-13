package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type FindPetsByCityUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewFindPetByFindPetsByCityUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *FindPetsByCityUseCase {
	return &FindPetsByCityUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *FindPetsByCityUseCase) Execute(city string) (dto.PetOutputDTO, error) {
	return dto.PetOutputDTO{}, nil
}
