package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type PetHandler struct {
	CreatePetUseCase *usecase.CreatePetUseCase
}

func NewPetHandler(createPetUseCase *usecase.CreatePetUseCase) *PetHandler {
	return &PetHandler{
		CreatePetUseCase: createPetUseCase,
	}
}

func (h *PetHandler) Create(c *gin.Context) {
	var input dto.PetInputDTO
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	output, err := h.CreatePetUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot register pet",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)
}
