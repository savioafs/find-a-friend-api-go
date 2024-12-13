package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type UpdatePetUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewUpdatePetUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *UpdatePetUseCase {
	return &UpdatePetUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *UpdatePetUseCase) Execute(id string) (dto.PetOutputDTO, error) {
	return dto.PetOutputDTO{}, nil
}
