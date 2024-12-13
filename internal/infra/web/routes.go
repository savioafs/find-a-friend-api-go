package web

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/database"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/web/handlers"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	petRepository := database.NewPetRepositoryPG(dbConn)
	organizationRepository := database.NewOrganizationRepository(dbConn)

	registerPetRoutes(router, petRepository, organizationRepository)
	// registerOrganizationRoutes(router, organizationRepository)
}

func registerPetRoutes(router *gin.Engine, petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) {
	createPetUseCase := petUseCase.NewCreatePetUseCase(petRepo, orgRepo)
	getPetByIDUseCase := petUseCase.NewFindPetByIDUseCase(petRepo, orgRepo)
	getPetByCityUseCase := petUseCase.NewFindPetByFindPetsByCityUseCase(petRepo, orgRepo)
	getAllPetsByCharacteristicsUseCase := petUseCase.NewAllPetsByCharacteristicsUseCase(petRepo, orgRepo)
	getAllPetsByOrganizationUseCase := petUseCase.NewAllPetsByOrganizationUseCase(petRepo, orgRepo)
	updatePet := petUseCase.NewUpdatePetUseCase(petRepo, orgRepo)
	deletePet := petUseCase.NewDeletePet(petRepo, orgRepo)

	petHandler := handlers.NewPetHandler(
		createPetUseCase,
		getPetByIDUseCase,
		getPetByCityUseCase,
		getAllPetsByCharacteristicsUseCase,
		getAllPetsByOrganizationUseCase,
		updatePet,
		deletePet)

	pet := router.Group("/pets")
	{
		// add middleware authenticator
		pet.POST("/", petHandler.Create)
		pet.GET("", petHandler.PetByID)
		pet.GET("", petHandler.PetsByCity)
		pet.GET("", petHandler.AllByCharacteristics)
		pet.GET("", petHandler.AllByOrganization)
		pet.PUT("", petHandler.Update)
		pet.DELETE("", petHandler.Delete)
	}
}
