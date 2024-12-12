package web

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/entity"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/database"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/web/handlers"
	organizationUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/organization"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	petRepository := database.NewPetRepositoryPG(dbConn)
	organizationRepository := database.NewOrganizationRepository(dbConn)

	registerPetRoutes(router, petRepository, organizationRepository)
	registerOrganizationRoutes(router, organizationRepository)
}

func registerPetRoutes(router *gin.Engine, petRepo entity.PetStorer, orgRepo entity.OrganizationStorer) {
	createPetUseCase := petUseCase.NewCreatePetUseCase(petRepo, orgRepo)
	petHandler := handlers.NewPetHandler(createPetUseCase)

	pet := router.Group("/pets")
	{
		// add middleware authenticator
		pet.POST("/", petHandler.Create)
		// pet.GET("", nil)    // find by id
		// pet.GET("", nil)    // find by city
		// pet.GET("", nil)    // find all by organization
		// pet.PUT("", nil)    // update pet
		// pet.DELETE("", nil) // delete pet
	}
}

func registerOrganizationRoutes(router *gin.Engine, orgRepo entity.OrganizationStorer) {
	organizationUseCase := organizationUseCase.NewCreateOrganizationUseCase(orgRepo)
	fmt.Println(organizationUseCase)
	router.Run("okok")
	// organization := router.Group("/organizations")
	// {
	// 	organization.POST("/", petHandler.Create)ß
	// 	organization.GET("", nil)    // find by id
	// 	organization.PUT("", nil)    // update org
	// 	organization.DELETE("", nil) // delete org

	// 	organization.POST("/generate_token", nil) // login
	// }

}
