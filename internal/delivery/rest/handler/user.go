package handler

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	body := domain.CreateUserReq{}

	c.ShouldBindJSON(&body)

	res, err := h.UserUseCase.CreateUser(domain.CreateUserReq{
		Name:  body.Name,
		Email: body.Email,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
