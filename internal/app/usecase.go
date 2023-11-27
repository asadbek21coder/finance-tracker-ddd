package app

import (
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UseCase struct {
	User domain.UserUseCase
}

func NewUseCase(db *sqlx.DB, repo domain.UserRepository) *UseCase {
	return &UseCase{
		User: NewUserUseCase(db, repo),
	}
}
