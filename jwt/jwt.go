package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

var sillyhatSigningKey = []byte("sillyhatisacleverboy")

func createTokenString(signingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(signingKey), claims)
	tokenstring, err := token.SignedString(sillyhatSigningKey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func CreateTokenStringHS256(claims jwt.Claims) (string, error) {
	return createTokenString("HS256", claims)
}

func CreateTokenStringHS384(claims jwt.Claims) (string, error) {
	return createTokenString("HS384", claims)
}

func CreateTokenStringHS512(claims jwt.Claims) (string, error) {
	return createTokenString("HS512", claims)
}

func ParseTokenString(tokenstring string, claims jwt.Claims) (*jwt.Token, error) {
	//token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
	//	return sillyhatSigningKey, nil
	//})
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return sillyhatSigningKey, nil
	})
	return token, err
}

//func ParseTokenStringHS256(claims jwt.Claims) (string, error) {
//	return parseTokenString("HS256", claims)
//}
//
//func ParseTokenStringHS384(claims jwt.Claims) (string, error) {
//	return createTokenString("HS384", claims)
//}
//
//func ParseTokenStringHS512(claims jwt.Claims) (string, error) {
//	return createTokenString("HS512", claims)
//}
