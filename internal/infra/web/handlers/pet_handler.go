package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/pet"
)

type PetHandler struct {
	CreatePetUseCase                   *usecase.CreatePetUseCase
	GetPetByIDUseCase                  *usecase.FindPetByIDUseCase
	GetPetByCityUseCase                *usecase.FindPetsByCityUseCase
	GetAllPetsByCharacteristicsUseCase *usecase.AllPetsByCharacteristicsUseCase
	GetAllPetsByOrganizationUseCase    *usecase.AllPetsByOrganizationUseCase
	UpdatePet                          *usecase.UpdatePetUseCase
	DeletePet                          *usecase.DeletePet
}

func NewPetHandler(
	createPetUseCase *usecase.CreatePetUseCase,
	getPetByIDUseCase *usecase.FindPetByIDUseCase,
	getPetByCityUseCase *usecase.FindPetsByCityUseCase,
	getAllPetsByCharacteristicsUseCase *usecase.AllPetsByCharacteristicsUseCase,
	getAllPetsByOrganizationUseCase *usecase.AllPetsByOrganizationUseCase,
	updatePet *usecase.UpdatePetUseCase,
	deletePet *usecase.DeletePet) *PetHandler {
	return &PetHandler{
		CreatePetUseCase:                   createPetUseCase,
		GetPetByIDUseCase:                  getPetByIDUseCase,
		GetPetByCityUseCase:                getPetByCityUseCase,
		GetAllPetsByCharacteristicsUseCase: getAllPetsByCharacteristicsUseCase,
		GetAllPetsByOrganizationUseCase:    getAllPetsByOrganizationUseCase,
		UpdatePet:                          updatePet,
		DeletePet:                          deletePet,
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

func (h *PetHandler) PetByID(c *gin.Context)              {}
func (h *PetHandler) PetsByCity(c *gin.Context)           {}
func (h *PetHandler) AllByCharacteristics(c *gin.Context) {}
func (h *PetHandler) AllByOrganization(c *gin.Context)    {}
func (h *PetHandler) Update(c *gin.Context)               {}
func (h *PetHandler) Delete(c *gin.Context)               {}
