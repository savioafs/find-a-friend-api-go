package web

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/database"
	organizationHandler "github.com/savioafs/find-a-friend-api-go/internal/infra/web/handler/organization"
	petHandler "github.com/savioafs/find-a-friend-api-go/internal/infra/web/handler/pet"
	organizationUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/organization"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	petRepository := database.NewPetRepositoryPG(dbConn)
	organizationRepository := database.NewOrganizationRepository(dbConn)

	// resolver isso
	var (
		jwtAuth  *jwtauth.JWTAuth
		expireIn int
	)

	registerPetRoutes(router, petRepository, organizationRepository)
	registerOrganizationRoutes(router, organizationRepository, jwtAuth, expireIn)

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

	pet := router.Group("/pet")
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

func registerOrganizationRoutes(router *gin.Engine, orgRepo entity.OrganizationStorer, jwtAuth *jwtauth.JWTAuth, expireIn int) {
	getJwtOrgUsecase := organizationUseCase.NewFindOrganizationByEmailJWT(orgRepo)
	getJwtHanlder := organizationHandler.NewGetJWTOrganizationHandler(getJwtOrgUsecase, jwtAuth, expireIn)

	organization := router.Group("/organization")
	{
		organization.POST("/generate_token", getJwtHanlder.Handle)
	}
}
