package handler

import (
	"github.com/gin-gonic/gin"
	petUseCase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type FindPetByIDHandler struct {
	FindPetByIDUseCase *petUseCase.FindPetByIDUseCase
}

func NewFindPetByIDHandler(findPetByIDUseCase *petUseCase.FindPetByIDUseCase) *FindPetByIDHandler {
	return &FindPetByIDHandler{
		FindPetByIDUseCase: findPetByIDUseCase,
	}
}

func (h *FindPetByIDHandler) Handle(c *gin.Context) {}
