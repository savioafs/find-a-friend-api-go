package web

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/database"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/web/handlers"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	petRepository := database.NewPetRepositoryPG(dbConn)
	organizationRepository := database.NewOrganizationRepository(dbConn)

	createPetUseCase := usecase.NewCreatePetUseCase(petRepository, organizationRepository)
	petHandler := handlers.NewPetHandler(createPetUseCase)

	// organization := router.Group("/organization")
	// {
	// 	organization.POST("/", petHandler.Create)ß
	// 	organization.GET("", nil)    // find by id
	// 	organization.PUT("", nil)    // update org
	// 	organization.DELETE("", nil) // delete org

	// 	organization.POST("/generate_token", nil) // login
	// }

	pet := router.Group("/pet")
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
