package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UserToken struct {
	UserId   string
	UserName string
	jwt.StandardClaims
}

func TestCreateTokenStringHS256(t *testing.T) {
	userToken := *&UserToken{UserId: "1001", UserName: "Cookie"}
	token, err := CreateTokenStringHS256(userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAxIiwiVXNlck5hbWUiOiJDb29raWUifQ.Tx4ENWMmpsvw9XQtt5hALOTJY-KmC0e-TpRDroivePw")
}

func TestCreateTokenStringHS384(t *testing.T) {
	userToken := *&UserToken{UserId: "1001", UserName: "Cookie"}
	token, err := CreateTokenStringHS384(userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAxIiwiVXNlck5hbWUiOiJDb29raWUifQ.XACwFIbPYehKI-F8zC3l_O3910hfnDcjT3ehNvH6OZXMOHimU8Ap6E50ZbXLplrM")
}

func TestCreateTokenStringHS512(t *testing.T) {
	userToken := *&UserToken{UserId: "1001", UserName: "Cookie"}
	token, err := CreateTokenStringHS512(userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAxIiwiVXNlck5hbWUiOiJDb29raWUifQ.LeBxRT_Kn9KuwL9rqpRbtL6vuN7x9kTckDOcde7RGeo4SR-_BEFnMnQsOYDWBxVS3sVoE-5Z2DTG7QF1aloaKA")
}

func TestParseTokenString(t *testing.T) {
	userToken := UserToken{}
	tokenstring := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAxIiwiVXNlck5hbWUiOiJDb29raWUifQ.Tx4ENWMmpsvw9XQtt5hALOTJY-KmC0e-TpRDroivePw"
	token, err := ParseTokenString(tokenstring, &userToken)
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.EqualValues(t, userToken.UserId, "1001")
	assert.EqualValues(t, userToken.UserName, "Cookie")
}
