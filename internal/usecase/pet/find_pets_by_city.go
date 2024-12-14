package usecase

import (
	"errors"

	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
)

type FindPetsByCityUseCase struct {
	PetRepository entity.PetStorer
	OrgRepository entity.OrganizationStorer
}

func NewFindPetByFindPetsByCityUseCase(petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) *FindPetsByCityUseCase {
	return &FindPetsByCityUseCase{
		PetRepository: petRepo,
		OrgRepository: orgRepo,
	}
}

func (u *FindPetsByCityUseCase) Execute(city string) ([]dto.PetOutputDTO, error) {

	pets, err := u.PetRepository.FindByCity(city)
	if err != nil {
		return nil, err
	}

	if len(pets) == 0 {
		return nil, errors.New("no pets found for the specified city")
	}

	org, err := u.OrgRepository.FindOrganizationByID(pets[0].OrgID)
	if err != nil {
		return nil, err
	}

	var petsByCity []dto.PetOutputDTO

	for _, pet := range pets {

		petOutputByCity := dto.PetOutputDTO{
			ID:                      pet.ID,
			Org:                     org.Name,
			OrgWhatspp:              org.Whatsapp,
			Name:                    pet.Name,
			About:                   pet.About,
			Age:                     pet.Age,
			Size:                    pet.Size,
			EnergyLevel:             pet.EnergyLevel,
			DependencyLevel:         pet.DependencyLevel,
			Environment:             pet.Environment,
			Photos:                  pet.Photos,
			RequirementsForAdoption: pet.RequirementsForAdoption,
		}

		petsByCity = append(petsByCity, petOutputByCity)
	}

	return petsByCity, nil
}
