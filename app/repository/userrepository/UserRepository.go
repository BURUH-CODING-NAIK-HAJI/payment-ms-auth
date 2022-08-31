package userrepository

import "github.com/rizface/golang-api-template/app/entity/responseentity"

type UserRepositoryInterface interface {
	FindOne(username string) (*responseentity.User, error)
	Create()
}

type UserRepository struct {
}

func New() UserRepositoryInterface {
	return &UserRepository{}
}

func (userrepository *UserRepository) FindOne(username string) (*responseentity.User, error) {
	return nil, nil
}

func (userrepository *UserRepository) Create() {
}
