package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
	"github.com/test-go/testify/assert"
)

func TestCreatePet(t *testing.T) {
	as := assert.New(t)

	db, mock, err := sqlmock.New()
	as.NoError(err)
	defer db.Close()

	pet, err := entity.NewPet(
		uuid.New().String(),
		"Zezinho Cachorrão",
		"Gente fina",
		"Filhote",
		"Medio",
		"Muito eletrico",
		"Muito Carente",
		"Amplo",
		[]string{"photo1.jpg", "photo2.jpg"},
		[]string{"espaço amplo", "ambiente saudavel"})
	as.NoError(err)
	as.NotNil(pet)

	query := "INSERT INTO pets \\(id, org_id, name, about, age, size, energy_level, dependency_level, enviroment, photos, requiriments_for_adoption\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11\\) RETURNING id"
	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(
			pet.ID, pet.OrgID, pet.Name, pet.About, pet.Age,
			pet.Size, pet.EnergyLevel, pet.DependencyLevel,
			pet.Environment, pq.Array(pet.Photos), pq.Array(pet.RequirementsForAdoption),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(pet.ID))

	repo := NewPetRepositoryPG(db)

	err = repo.Save(pet)
	as.NoError(err)

	err = mock.ExpectationsWereMet()
	as.NoError(err)
}
