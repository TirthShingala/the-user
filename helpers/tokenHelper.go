package helper

import (
	"errors"
	"log"
	"time"

	"github.com/TirthShingala/the-user/constants"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	Role  string
	jwt.StandardClaims
}

var SECRET_KEY string = constants.EnvSecretKey()

func GenerateAllTokens(email string, role string) (signedToken string, signedRefreshToken string, err error) {
	token, _ := GenerateAccessTokens(email, role)

	refreshClaims := &SignedDetails{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(constants.RefreshTokenDurationInHour)).Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func GenerateAccessTokens(email string, role string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(constants.TokenDurationInHour)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		log.Println(err.Error())
		err = errors.New("the token is invalid")
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err = errors.New("the token is invalid")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return nil, err
	}
	return claims, nil
}
