package usecase

import (
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
	return dto.PetOutputDTO{}, nil
}
