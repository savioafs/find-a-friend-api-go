package handler

import (
	"github.com/gin-gonic/gin"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type FindPetsByCityHandler struct {
	FindPetsByCityUseCase *petUseCase.FindPetsByCityUseCase
}

func NewFindPetByCityHandler(findPetsByCityUseCase *petUseCase.FindPetsByCityUseCase) *FindPetsByCityHandler {
	return &FindPetsByCityHandler{
		FindPetsByCityUseCase: findPetsByCityUseCase,
	}
}

func (h *FindPetsByCityHandler) Handle(c *gin.Context) {}
