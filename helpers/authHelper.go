package helper

import (
	"errors"
	"log"

	"github.com/TirthShingala/the-user/constants"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var PASSWORD_PEPPER = constants.EnvPasswordPepper()

func CheckUserType(ctx *gin.Context, role string) (err error) {
	userRole := ctx.GetString("role")
	err = nil
	if userRole != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+PASSWORD_PEPPER), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword+PASSWORD_PEPPER))
	check := true

	if err != nil {
		log.Println(err)
		check = false
	}

	return check
}
