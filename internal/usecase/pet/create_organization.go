package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type OrganizationUseCase struct {
	OrganizationRepository entity.OrganizationStorer
}

func NewOrganizationRepository(repo entity.OrganizationStorer) *OrganizationUseCase {
	return &OrganizationUseCase{
		OrganizationRepository: repo,
	}
}
