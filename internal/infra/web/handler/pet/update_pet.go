package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type UpdatePetHandler struct {
	UpdatePetUseCase *usecase.UpdatePetUseCase
}

func NewUpdatePetHandler(updatePetUseCase *usecase.UpdatePetUseCase) *UpdatePetHandler {
	return &UpdatePetHandler{
		UpdatePetUseCase: updatePetUseCase,
	}
}

func (h *UpdatePetHandler) Handle(c *gin.Context) {}
