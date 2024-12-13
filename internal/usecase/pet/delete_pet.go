package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type DeletePet struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewDeletePet(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *DeletePet {
	return &DeletePet{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *DeletePet) Execute(id string) (dto.PetOutputDTO, error) {
	return dto.PetOutputDTO{}, nil
}
