package handler

import (
	"net/http"

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

func (h *FindPetByIDHandler) Handle(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message:": "id cannot empty",
		})
		return
	}

	pet, err := h.FindPetByIDUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pet)
}
