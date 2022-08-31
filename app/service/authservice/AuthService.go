package authservice

import (
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/securityentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"github.com/rizface/golang-api-template/system/security"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(payload *requestentity.Login) securityentity.GeneratedResponseJwt
	Register()
}

type AuthService struct {
	userrepository userrepository.UserRepositoryInterface
}

func New(userrepository userrepository.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		userrepository: userrepository,
	}
}

func (authservice *AuthService) Login(payload *requestentity.Login) securityentity.GeneratedResponseJwt {
	existingUser, err := authservice.userrepository.FindOne(payload.Username)
	if err != nil {
		panic(errorgroup.InternalServerError)
	}
	if existingUser == nil {
		panic(errorgroup.USER_NOT_FOUND)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(payload.Password))
	if err != nil {
		panic(errorgroup.UNAUTHORIZED)
	}

	userDataForGenerateToken := &securityentity.UserData{
		Id:       existingUser.Id,
		Name:     existingUser.Name,
		Username: existingUser.Username,
	}
	generatedTokenSchema := security.GenerateToken(userDataForGenerateToken)
	return generatedTokenSchema
}

func (authservice *AuthService) Register() {
}
