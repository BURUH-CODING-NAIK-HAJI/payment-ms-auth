package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rizface/golang-api-template/app/entity/securityentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/config"
)

func getSecret() (string, string) {
	if len(os.Getenv("JWT_BEARER_SECRET")) == 0 {
		return config.JWT_BEARER_SECRET, config.JWT_REFRESH_SECRET
	}
	return os.Getenv("JWT_BEARER_SECRET"), os.Getenv("JWT_REFRESH_SECRET")
}

type JwtClaim struct {
	userData securityentity.UserData
	jwt.RegisteredClaims
}

func EncodeDataToToken(secret string, userData *securityentity.UserData) string {
	claim := JwtClaim{
		*userData,
		jwt.RegisteredClaims{
			Issuer:    config.JWT_ISSUER,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   userData.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWT_EXPIRE * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf(err.Error())
	}

	return tokenString
}

func GenerateToken(userData *securityentity.UserData) securityentity.GeneratedResponseJwt {
	bearerSecret, refreshSecret := getSecret()
	bearerToken := EncodeDataToToken(bearerSecret, userData)
	refreshToken := EncodeDataToToken(refreshSecret, userData)

	return securityentity.GeneratedResponseJwt{
		UserData: *userData,
		TokenSchema: securityentity.TokenSchema{
			Bearer:  bearerToken,
			Refresh: refreshToken,
		},
	}
}

func DecodeToken(tokenString, tokenType string) JwtClaim {
	var secret string
	if tokenType == "bearer" {
		secret = config.JWT_BEARER_SECRET
	} else {
		secret = config.JWT_REFRESH_SECRET
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			panic(errorgroup.TOKEN_EXPIRED)
		} else if errors.Is(err, jwt.ErrSignatureInvalid) {
			panic(errorgroup.TOKEN_INVALID)
		} else {
			panic(errorgroup.UNAUTHORIZED)
		}
	}
	if claims, ok := token.Claims.(*JwtClaim); ok && token.Valid {
		return *claims
	} else {
		panic(errorgroup.UNAUTHORIZED)
	}
}
