package database

import (
	"database/sql"
	"errors"

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

	err = stmt.QueryRow(pet.ID, pet.OrgID, pet.Name, pet.About, pet.Age, pet.Size, pet.EnergyLevel, pet.DependencyLevel, pet.Environment, pq.Array(pet.Photos), pq.Array(pet.RequirementsForAdoption)).Err()
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
		return nil, errors.New("pet does not exist")
	}

	return &pet, nil
}

func (r *PetRepositoryPG) FindByCity(city string) ([]entity.Pet, error) {
	query := "SELECT p.* FROM pets p JOIN organizations o ON p.org_id = o.id WHERE o.city = $1"

	rows, err := r.DB.Query(query, city)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pets []entity.Pet

	for rows.Next() {
		var pet entity.Pet

		err := rows.Scan(
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

		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *PetRepositoryPG) AllByOrganization(org_id string) ([]entity.Pet, error) {
	query := "SELECT p.* FROM pet p JOIN organization o ON p.organization_id = o.id WHERE o.id = $1"

	rows, err := r.DB.Query(query, org_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pets []entity.Pet

	for rows.Next() {
		var pet entity.Pet

		err := rows.Scan(
			&pet.ID,
			&pet.OrgID,
			&pet.Name,
			&pet.About,
			&pet.Age,
			&pet.Size,
			&pet.EnergyLevel,
			&pet.DependencyLevel,
			&pet.Environment,
			&pet.Photos,
			&pet.RequirementsForAdoption)

		if err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *PetRepositoryPG) AllByCharacteristics(characteristics []string) ([]entity.Pet, error) {
	query := "SELECT * FROM pet WHERE characteristics && $1"

	rows, err := r.DB.Query(query, pq.Array(characteristics))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pets []entity.Pet

	for rows.Next() {
		var pet entity.Pet

		err := rows.Scan(
			&pet.ID,
			&pet.OrgID,
			&pet.Name,
			&pet.About,
			&pet.Age,
			&pet.Size,
			&pet.EnergyLevel,
			&pet.DependencyLevel,
			&pet.Environment,
			&pet.Photos,
			&pet.RequirementsForAdoption)

		if err != nil {
			return nil, err
		}

		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *PetRepositoryPG) Update(id string, pet *entity.Pet) (*entity.Pet, error) {
	petFind, err := r.FindPetByID(id)
	if err != nil {
		return nil, err
	}

	if petFind == nil {
		return nil, err
	}

	query := `UPDATE pets 
		SET name = $1, about = $2, age = $3, size = $4, energy_level = $5, dependency_level = $6, environment = $7, photos = $8, requirements_for_adoption = $9
		WHERE id = $10
		RETURNING id, org_id, name, about, age, size, energy_level, dependency_level, environment, photos, requirements_for_adoption`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var updatedPet entity.Pet
	err = stmt.QueryRow(
		pet.Name,
		pet.About,
		pet.Age,
		pet.Size,
		pet.EnergyLevel,
		pet.DependencyLevel,
		pet.Environment,
		pq.Array(pet.Photos),
		pq.Array(pet.RequirementsForAdoption),
		id,
	).Scan(
		&updatedPet.ID,
		&updatedPet.OrgID,
		&updatedPet.Name,
		&updatedPet.About,
		&updatedPet.Age,
		&updatedPet.Size,
		&updatedPet.EnergyLevel,
		&updatedPet.DependencyLevel,
		&updatedPet.Environment,
		pq.Array(&updatedPet.Photos),
		pq.Array(&updatedPet.RequirementsForAdoption),
	)
	if err != nil {
		return nil, err
	}

	return &updatedPet, nil
}

func (r *PetRepositoryPG) Delete(id string) error {
	pet, err := r.FindPetByID(id)
	if err != nil {
		return err
	}

	if pet == nil {
		return errors.New("no pet found with the given id")
	}

	query := "DELETE FROM pets WHERE id = $1"

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no pet found with the given id")
	}

	return nil
}
