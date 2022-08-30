package security_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rizface/golang-api-template/app/entity/securityentity"
	"github.com/rizface/golang-api-template/system/security"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDataToJwt(t *testing.T) {
	userData := &securityentity.UserData{
		Id:   "id",
		Name: "Fariz",
	}
	generatedResponseJwt := security.GenerateToken(userData)
	assert.Equal(t, "securityentity.GeneratedResponseJwt", reflect.TypeOf(generatedResponseJwt).String())
	assert.Equal(t, "string", reflect.TypeOf(generatedResponseJwt.TokenSchema.Bearer).String())
	assert.Equal(t, "string", reflect.TypeOf(generatedResponseJwt.TokenSchema.Refresh).String())
}

func TestDecodeBearerJwt(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	claim := security.DecodeToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJSSVpGQUNFIiwic3ViIjoiRmFyaXoiLCJleHAiOjE2NjE3MDQ4NjAsIm5iZiI6MTY2MTUzMjA2MCwiaWF0IjoxNjYxNTMyMDYwfQ.gg84YYMtjXwqIpedGjnTVJ2bwdleSPOIHxe61SrNt_0", "bearer")
	assert.Equal(t, "security.JwtClaim", reflect.TypeOf(claim).String())
}

func TestDecodeRefreshToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJSSVpGQUNFIiwic3ViIjoiRmFyaXoiLCJleHAiOjE2NjE3MDUwMjgsIm5iZiI6MTY2MTUzMjIyOCwiaWF0IjoxNjYxNTMyMjI4fQ.bN-rqM0qPENTnReTDpS2qh8jLWO651jBzewfiOomUM8"
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	claim := security.DecodeToken(token, "refresh")
	assert.Equal(t, "security.JwtClaim", reflect.TypeOf(claim).String())
}
