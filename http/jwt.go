package http

import (
	"github.com/Barna001/go-auth/errors"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Email   string `json:"email"`
	Methods string `json:"methods"`
	jwt.StandardClaims
}

func CreateTokenForUsersEndpoint(signingKey string, email string) string {
	claims := UserClaims{
		email,
		"GET, POST",
		jwt.StandardClaims{
			ExpiresAt: 10000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	errors.CriticalHandling(err)
	return signedToken
}

func GetClaimsFromToken(tokenString string, signingKey string) (string, error) {
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
		return claims.Methods, nil
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
