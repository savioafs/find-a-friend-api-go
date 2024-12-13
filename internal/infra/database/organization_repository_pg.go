package database

import (
	"database/sql"

	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type OrganizationRepository struct {
	DB *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
	return &OrganizationRepository{
		DB: db,
	}
}

func (r *OrganizationRepository) FindOrganizationByID(id string) (*entity.Organization, error) {
	stmt, err := r.DB.Prepare("SELECT id, name, owner, email, cep, city, whatsapp FROM organizations WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var organization entity.Organization

	err = stmt.QueryRow(id).Scan(
		&organization.ID,
		&organization.Name,
		&organization.Owner,
		&organization.Email,
		&organization.CEP,
		&organization.City,
		&organization.Whatsapp)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *OrganizationRepository) FindOrganizationByEmail(email string) (*entity.Organization, error) {
	return nil, nil
}
func (r *OrganizationRepository) Save(org *entity.Organization) error { return nil }
func (r *OrganizationRepository) UpdateOrganization(id string) (*entity.Organization, error) {
	return nil, nil
}
func (r *OrganizationRepository) DeleteOrganization(id string) error { return nil }
