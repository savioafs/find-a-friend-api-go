package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type CreateOrganizationUseCase struct {
	OrganizationRepository entity.OrganizationStorer
}

func NewCreateOrganizationUseCase(repo entity.OrganizationStorer) *CreateOrganizationUseCase {
	return &CreateOrganizationUseCase{
		OrganizationRepository: repo,
	}
}

func (u *CreateOrganizationUseCase) Execute(input dto.OrganizationInputDTO) (dto.OrganizationOutputDTO, error) {
	return dto.OrganizationOutputDTO{}, nil
}
