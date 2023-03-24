package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/TirthShingala/the-user/constants"
	helper "github.com/TirthShingala/the-user/helpers"
	"github.com/TirthShingala/the-user/models"
	"github.com/TirthShingala/the-user/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user models.User

		if err := ctx.BindJSON(&user); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "invalid json!",
			})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			log.Println(validationErr.Error())
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: validationErr.Error(),
			})
			return
		}
		if user.Role == "" {
			user.Role = constants.UserRole
		}
		if user.Role != constants.UserRole && user.Role != constants.AdminRole {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "role must be USER or ADMIN!",
			})
			return
		}

		count, err := models.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "this email already exists",
			})
			return
		}

		count, err = models.UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "this phone number already exists",
			})
			return
		}

		password, err := helper.HashPassword(user.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		user.Password = password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.IsBlocked = false
		token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.Role)

		_, insertErr := models.UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "user is not created",
			})
			return
		}

		ctx.JSON(http.StatusOK, responses.AuthResponse{
			Success:      true,
			Message:      "user created successfully",
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Phone:        user.Phone,
			Role:         user.Role,
			IsBlocked:    user.IsBlocked,
			Token:        token,
			RefreshToken: refreshToken,
		})
	}
}

type loginReq struct {
	Email    string
	Password string
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req loginReq

		if err := ctx.BindJSON(&req); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "invalid json!",
			})
			return
		}

		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "email not provided!",
			})
			return
		}

		if req.Password == "" {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "password not provided!",
			})
			return
		}

		var user models.User

		err := models.UserCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "email or password is incorrect!",
			})
			return
		}

		if user.IsBlocked {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: "you are blocked!",
			})
			return
		}

		passwordIsValid := helper.VerifyPassword(user.Password, req.Password)
		if !passwordIsValid {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "email or password is incorrect!",
			})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.Role)

		ctx.JSON(http.StatusOK, responses.AuthResponse{
			Success:      true,
			Message:      "user logged in successfully",
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Phone:        user.Phone,
			Role:         user.Role,
			IsBlocked:    user.IsBlocked,
			Token:        token,
			RefreshToken: refreshToken,
		})
	}
}

func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "token not provided!",
			})
			ctx.Abort()
			return
		}

		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "Bearer token not in proper format!",
			})
			ctx.Abort()
			return
		}

		token = splitToken[1]

		claims, err := helper.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		var user models.User

		mongoerr := models.UserCollection.FindOne(ctx, bson.M{"email": claims.Email}).Decode(&user)
		if mongoerr != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "email incorrect!",
			})
			return
		}

		if user.IsBlocked {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: "you are blocked!",
			})
			return
		}

		newtoken, _ := helper.GenerateAccessTokens(user.Email, user.Role)

		ctx.JSON(http.StatusOK, responses.AuthResponse{
			Success:   true,
			Message:   "new access token generated successfully",
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Role:      user.Role,
			IsBlocked: user.IsBlocked,
			Token:     newtoken,
		})
	}
}
