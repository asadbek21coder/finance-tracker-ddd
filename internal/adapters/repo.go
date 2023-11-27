package adapters

import (
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	User domain.UserRepository
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		User: NewUserRepository(db),
	}
}
