package database

import (
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type PetRepositoryPG struct {
	DB *sql.DB
}

func NewPetRepositoryPG(db *sql.DB) *PetRepositoryPG {
	return &PetRepositoryPG{DB: db}
}

func (r *PetRepositoryPG) Save(pet *entity.Pet) error {

	stmt, err := r.DB.Prepare("INSERT INTO pets (id, org_id, name, about, age, size, energy_level, dependency_level, enviroment, photos, requirements_for_adoption) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id")
	if err != nil {
		return err
	}

	defer stmt.Close()

	photosJSON, err := json.Marshal(pet.Photos)
	if err != nil {
		return err
	}

	requirementsJSON, err := json.Marshal(pet.RequirementsForAdoption)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(pet.ID, pet.OrgID, pet.Name, pet.About, pet.Age, pet.Size, pet.EnergyLevel, pet.DependencyLevel, pet.Environment, photosJSON, requirementsJSON).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *PetRepositoryPG) FindPetByID(id string) (*entity.Pet, error) {
	stmt, err := r.DB.Prepare("SELECT id, org_id, name, about, age, size, energy_level, dependency_level, enviroment, photos, requirements_for_adoption FROM pets WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var pet entity.Pet

	err = stmt.QueryRow(id).Scan(
		&pet.ID,
		&pet.OrgID,
		&pet.Name,
		&pet.About,
		&pet.Age,
		&pet.Size,
		&pet.EnergyLevel,
		&pet.DependencyLevel,
		&pet.Environment,
		pq.Array(&pet.Photos),
		pq.Array(&pet.RequirementsForAdoption))
	if err != nil {
		return nil, err
	}

	return &pet, nil
}
