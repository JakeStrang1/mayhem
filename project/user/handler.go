package user

import (
	"example.com/myNetwork/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	controller ControllerInterface
}

func NewHandler(controller ControllerInterface) *Handler {
	return &Handler{controller: controller}
}

type ControllerInterface interface {
	SignUp(string) (*models.User, error)
}

func (h *Handler) SignUp(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "malformed request"})
		return
	}

	v := validator.New()
	v.SetTagName("sign_up")
	err = v.Struct(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
		return
	}

	dbUser, err := h.controller.SignUp(user.Email)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
		return
	}
	user = FromDBUser(*dbUser)
	c.JSON(200, user)
}
