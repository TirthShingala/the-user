package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/TirthShingala/the-user/constants"
	helper "github.com/TirthShingala/the-user/helpers"
	"github.com/TirthShingala/the-user/models"
	"github.com/TirthShingala/the-user/responses"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if authErr := helper.CheckUserType(ctx, constants.AdminRole); authErr != nil {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: authErr.Error(),
			})
			return
		}

		limitQuery := ctx.Query("limit")
		pageQuery := ctx.Query("page")

		limit, _ := strconv.ParseInt(limitQuery, 0, 64)
		page, _ := strconv.ParseInt(pageQuery, 0, 64)
		if limit == 0 {
			limit = 20
		}
		if limit > 50 {
			limit = 50
		}

		skip := int64(page * limit)
		fOpt := options.FindOptions{Limit: &limit, Skip: &skip}

		curr, err := models.UserCollection.Find(ctx, bson.M{}, &fOpt)
		if err != nil {
			log.Panic(err.Error())
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		users := make([]responses.AdminUserResponse, 0)

		for curr.Next(ctx) {
			var user responses.AdminUserResponse
			if err := curr.Decode(&user); err != nil {
				log.Println(err)
			}
			users = append(users, user)
		}

		count, _ := models.UserCollection.CountDocuments(ctx, bson.M{})

		metaData := responses.PaginationMetaData{
			Page:       page,
			PerPage:    limit,
			PageCount:  count / limit,
			TotalCount: count,
		}

		ctx.JSON(http.StatusOK, responses.PaginationResponse{
			Success:  true,
			Message:  "users fetched successfully",
			MetaData: metaData,
			Data:     users,
		})
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
			Message:   "profile data found successfully",
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Role:      user.Role,
			IsBlocked: user.IsBlocked,
		})
	}
}

type uploadUrlReq struct {
	ContentLength int64
	Extension     string
}

func UploadURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("email")

		var req uploadUrlReq
		if err := ctx.BindJSON(&req); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "invalid json!",
			})
			return
		}

		if req.ContentLength > 5242880 {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "file greater than 5MB is not allowed!",
			})
			return
		}

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

		key := "profile/" + user.ID.Hex() + "." + req.Extension

		signedPutUrl := helper.SignedURL(&key, &req.ContentLength)

		getUrl := constants.CDN_BASE_URL + "/" + key

		ctx.JSON(http.StatusOK, responses.SignedUrlResponse{
			Success:      true,
			Message:      "signed url created successfully",
			SignedPutUrl: signedPutUrl,
			Url:          getUrl,
		})
	}
}
