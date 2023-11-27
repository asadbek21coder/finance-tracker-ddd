package rest

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	body := domain.CreateUserReq{}

	c.ShouldBindJSON(&body)

	res, err := h.usecase.CreateUser(domain.CreateUserReq{
		Name:  body.Name,
		Email: body.Email,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
