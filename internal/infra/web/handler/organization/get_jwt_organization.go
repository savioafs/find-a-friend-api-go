package hanlder

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/savioafs/find-a-friend-api-go/internal/dto"
	usecase "github.com/savioafs/find-a-friend-api-go/internal/usecase/organization"
)

type GetJWTOrganizationHandler struct {
	FindOrganizationByEmailJWT *usecase.FindOrganizationByEmailJWT
	Jwt                        *jwtauth.JWTAuth
	JwtExpiresIn               int
}

func NewGetJWTOrganizationHandler(orgJwtFindByEmail *usecase.FindOrganizationByEmailJWT, jwt *jwtauth.JWTAuth, expiresJwt int) *GetJWTOrganizationHandler {
	return &GetJWTOrganizationHandler{
		FindOrganizationByEmailJWT: orgJwtFindByEmail,
		Jwt:                        jwt,
		JwtExpiresIn:               expiresJwt,
	}
}

func (h *GetJWTOrganizationHandler) Handle(c *gin.Context) {
	var org dto.GetJWTInput
	err := c.BindJSON(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message:": "invalid input to login",
			"error:":   err.Error(),
		})
		return
	}

	orgFind, err := h.FindOrganizationByEmailJWT.Execute(org.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message:": "invalid data to login",
			"error:":   err.Error(),
		})
		return
	}

	if !orgFind.ValidatePassword(org.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message:": "email or password invalid",
		})
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": orgFind.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	c.JSON(http.StatusOK, accessToken)
}
