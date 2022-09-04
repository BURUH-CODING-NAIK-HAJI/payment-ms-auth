package userrepository

import (
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
)

type UserRepositoryInterface interface {
	FindOne(username string) (*responseentity.User, error)
	Create(payload *requestentity.Register) (*responseentity.User, error)
}

type UserRepository struct {
}

func New() UserRepositoryInterface {
	return &UserRepository{}
}

func (userrepository *UserRepository) FindOne(username string) (*responseentity.User, error) {
	return nil, nil
}

func (userrepository *UserRepository) Create(payload *requestentity.Register) (*responseentity.User, error) {
	return nil, nil
}
