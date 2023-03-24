package controllers

import (
	"log"
	"net/http"

	"github.com/TirthShingala/the-user/constants"
	helper "github.com/TirthShingala/the-user/helpers"
	"github.com/TirthShingala/the-user/models"
	"github.com/TirthShingala/the-user/responses"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.CheckUserType(ctx, constants.AdminRole); err != nil {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("email")

		var user models.User
		err := models.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		if err != nil {
			log.Panic(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		ctx.JSON(http.StatusOK, responses.UserResponse{
			Success:   true,
			Message:   "user data found successfully",
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Role:      user.Role,
			IsBlocked: user.IsBlocked,
		})
	}
}
