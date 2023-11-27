package domain

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	CreateUser(CreateUserReq) (User, error)
}

type UserUseCase interface {
	CreateUser(CreateUserReq) (User, error)
}
