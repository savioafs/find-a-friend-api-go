package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type FindPetsByCharacteristicsHandler struct {
	FindPetsByCharacteristicsUseCase *usecase.AllPetsByCharacteristicsUseCase
}

func NewFindPetsByCharacteristicsHandler(findPetsByCharacteristicsUseCase *usecase.AllPetsByCharacteristicsUseCase) *FindPetsByCharacteristicsHandler {
	return &FindPetsByCharacteristicsHandler{
		FindPetsByCharacteristicsUseCase: findPetsByCharacteristicsUseCase,
	}
}

func (h *FindPetsByCharacteristicsHandler) Handle(c *gin.Context) {}
