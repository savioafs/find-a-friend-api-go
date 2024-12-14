package handler

import (
	"net/http"

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

func (h *FindPetsByCityHandler) Handle(c *gin.Context) {
	city := c.Param("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message:": "city cannot empty",
		})
		return
	}

	pets, err := h.FindPetsByCityUseCase.Execute(city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pets)
}
