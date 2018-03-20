package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/Barna001/go-auth/errors"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Email           string `json:"email"`
	EndpointMethods string `json:"methods"`
	jwt.StandardClaims
}

func createTokenForEndpoints(signingKey string, email string) string {
	claims := UserClaims{
		email,
		"user/GET, user/POST",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60, // 60 sec exp
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(signingKey))
	errors.CriticalHandling(err)
	return signedToken
}

func getClaimsFromToken(tokenString string, signingKey string) (string, error) {
	if tokenString == "" {
		return "", errors.UnparsableTokenError{Message: "No JWT token"}
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if token.Valid {
		return getClaimsFromValidToken(token)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		return getValidationErrors(ve)
	}
	return "", errors.UnparsableTokenError{}
}

func getClaimsFromValidToken(token *jwt.Token) (string, error) {
	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims.EndpointMethods, nil
	} else {
		return "", errors.WrongTypeOfClaimsTokenError{Message: "Methods string and StandardClaims needed"}
	}
}

func getValidationErrors(ve *jwt.ValidationError) (string, error) {
	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		//Malformed token
		return "", errors.UnparsableTokenError{Message: ve.Error()}
	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		//Expired or not activated
		return "", errors.NotActivatedOrExpiredTokenError{Message: ve.Error()}
	}
	return "", errors.UnparsableTokenError{}
}

func getJwtTokenFromHeader(header http.Header) string {
	return strings.Replace(header.Get("Authentication"), "Bearer ", "", 1)
}
