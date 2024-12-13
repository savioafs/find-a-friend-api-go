package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type DeletePetHandler struct {
	DeletePetUseCase *usecase.DeletePet
}

func NewDeletePetHandler(deletePetUseCase *usecase.DeletePet) *DeletePetHandler {
	return &DeletePetHandler{
		DeletePetUseCase: deletePetUseCase,
	}
}

func (h *DeletePetHandler) Handle(c *gin.Context) {}
