package authservice

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/entity/securityentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"github.com/rizface/golang-api-template/system/security"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(payload *requestentity.Login) securityentity.GeneratedResponseJwt
	Register(payload *requestentity.Register) *responseentity.User
}

type AuthService struct {
	userrepository userrepository.UserRepositoryInterface
	redis          *redis.Client
}

func New(
	userrepository userrepository.UserRepositoryInterface,
	redis *redis.Client,
) AuthServiceInterface {
	return &AuthService{
		userrepository: userrepository,
		redis:          redis,
	}
}

func (authservice *AuthService) Login(payload *requestentity.Login) securityentity.GeneratedResponseJwt {
	existingUser, err := authservice.userrepository.FindOne(payload.Username)
	if err != nil {
		// NEED LOG ORIGINAL ERROR MESSAGE
		panic(errorgroup.InternalServerError)
	}
	if existingUser == nil {
		panic(errorgroup.USER_NOT_FOUND)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password.(string)), []byte(payload.Password))
	if err != nil {
		panic(errorgroup.UNAUTHORIZED)
	}

	userDataForGenerateToken := &securityentity.UserData{
		Id:       existingUser.Id,
		Name:     existingUser.Profile.Name,
		Username: existingUser.Username,
	}
	generatedTokenSchema := security.GenerateToken(userDataForGenerateToken)

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, time.Minute*1)
	defer cancel()
	bytesGeneratedTokenSchema, _ := json.Marshal(generatedTokenSchema)
	_, err = authservice.redis.Set(
		ctx,
		userDataForGenerateToken.Id,
		bytesGeneratedTokenSchema,
		24*time.Hour*30,
	).Result()
	if err != nil {
		panic(err)
	}
	return generatedTokenSchema
}

func (authservice *AuthService) Register(payload *requestentity.Register) *responseentity.User {
	result, err := authservice.userrepository.Create(payload)
	if err != nil {
		// NEED LOG ORIGINAL ERROR MESSAGE
		panic(errorgroup.REGISTER_FAILED)
	}
	return result
}
