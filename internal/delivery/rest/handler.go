package rest

import (
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/vertica/vertica-sql-go/logger"
)

type Handler struct {
	log     *logger.Logger
	usecase domain.UserUseCase
}

func NewHandler(usecase domain.UserUseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.pong)

	router.POST("/users", h.createUser)

	return router
}
