package app

import (
	"github.com/asadbek21coder/fintracker2/internal/adapters"
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(db *sqlx.DB) domain.UserUseCase {
	return UserUseCase{userRepo: adapters.NewUserRepository(db)}
}

func (u UserUseCase) CreateUser(user domain.CreateUserReq) (domain.User, error) {
	return u.userRepo.CreateUser(user)

}
