package adapters

import (
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(req domain.CreateUserReq) (domain.User, error) {
	res := domain.User{
		Id:    1,
		Name:  req.Name,
		Email: req.Email,
	}
	return res, nil
}
