package web

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/database"
	petHandler "github.com/savioafs/find-a-friend-api-go/internal/infra/web/handlers/pet"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	petRepository := database.NewPetRepositoryPG(dbConn)
	organizationRepository := database.NewOrganizationRepository(dbConn)

	registerPetRoutes(router, petRepository, organizationRepository)

}

func registerPetRoutes(router *gin.Engine, petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) {
	createPetUseCase := petUseCase.NewCreatePetUseCase(petRepo, orgRepo)
	createPetHandler := petHandler.NewCreatePetHandler(createPetUseCase)

	getPetByIDUseCase := petUseCase.NewFindPetByIDUseCase(petRepo, orgRepo)
	getPetByIDUseHandler := petHandler.NewFindPetByIDHandler(getPetByIDUseCase)

	getPetByCityUseCase := petUseCase.NewFindPetByFindPetsByCityUseCase(petRepo, orgRepo)
	getPetByCityHandler := petHandler.NewFindPetByCityHandler(getPetByCityUseCase)

	getAllPetsByCharacteristicsUseCase := petUseCase.NewAllPetsByCharacteristicsUseCase(petRepo, orgRepo)
	getAllPetsByCharacteristicsHandle := petHandler.NewFindPetsByCharacteristicsHandler(getAllPetsByCharacteristicsUseCase)

	getAllPetsByOrganizationUseCase := petUseCase.NewAllPetsByOrganizationUseCase(petRepo, orgRepo)
	getAllPetsByOrganizationHandler := petHandler.NewFindPetsByOrganizationHandler(getAllPetsByOrganizationUseCase)

	updatePetUseCase := petUseCase.NewUpdatePetUseCase(petRepo, orgRepo)
	updatePetHandler := petHandler.NewUpdatePetHandler(updatePetUseCase)

	deletePetUseCase := petUseCase.NewDeletePet(petRepo, orgRepo)
	deletePetHandler := petHandler.NewDeletePetHandler(deletePetUseCase)

	pet := router.Group("/pets")
	{
		// add middleware authenticator
		pet.POST("/", createPetHandler.Handle)
		pet.GET("/:id", getPetByIDUseHandler.Handle)
		pet.GET("/city/:city", getPetByCityHandler.Handle)
		pet.GET("/characteristics", getAllPetsByCharacteristicsHandle.Handle)
		pet.GET("/organization/:orgID", getAllPetsByOrganizationHandler.Handle)
		pet.PUT("/:id", updatePetHandler.Handle)
		pet.DELETE("/:id", deletePetHandler.Handle)
	}
}
