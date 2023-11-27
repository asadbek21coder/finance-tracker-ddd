package router

import (
	"github.com/asadbek21coder/fintracker2/internal/app"
	"github.com/asadbek21coder/fintracker2/internal/delivery/rest/handler"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB) *gin.Engine {
	router := gin.New()

	UserRoutes(router, db)
	return router
}

func UserRoutes(router *gin.Engine, db *sqlx.DB) {

	userHandler := handler.UserHandler{UserUseCase: app.NewUserUseCase(db)}

	users := router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
	}
}
