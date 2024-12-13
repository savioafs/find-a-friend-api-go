package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type FindPetsByOrganizationHandler struct {
	FindPetsByOrganizationUseCase *usecase.AllPetsByOrganizationUseCase
}

func NewFindPetsByOrganizationHandler(findPetsByOrganizationUseCase *usecase.AllPetsByOrganizationUseCase) *FindPetsByOrganizationHandler {
	return &FindPetsByOrganizationHandler{
		FindPetsByOrganizationUseCase: findPetsByOrganizationUseCase,
	}
}

func (h *FindPetsByOrganizationHandler) Handle(c *gin.Context) {}
