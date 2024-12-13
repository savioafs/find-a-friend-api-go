package usecase

import (
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type FindOrganizationByEmailJWT struct {
	OrgRepository entity.OrganizationStorer
}

func NewFindOrganizationByEmailJWT(orgRepo entity.OrganizationStorer) *FindOrganizationByEmailJWT {
	return &FindOrganizationByEmailJWT{
		OrgRepository: orgRepo,
	}
}

func (u *FindOrganizationByEmailJWT) Execute(email string) (entity.Organization, error) {
	return entity.Organization{}, nil
}
