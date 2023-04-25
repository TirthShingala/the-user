package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/TirthShingala/the-user/constants"
	helper "github.com/TirthShingala/the-user/helpers"
	"github.com/TirthShingala/the-user/models"
	"github.com/TirthShingala/the-user/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleClientID     = constants.EnvGoogleClientID()
	googleClientSecret = constants.EnvGoogleClientSecret()
	googleRedirectURI  = constants.EnvGoogleRedirectURI()
)

var googleOAuth2Config = &oauth2.Config{
	ClientID:     googleClientID,
	ClientSecret: googleClientSecret,
	RedirectURL:  googleRedirectURI,
	Endpoint:     google.Endpoint,
}

type googleOAuth2Response struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"Name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type googleReq struct {
	Code string
}

func Google() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req googleReq
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "invalid json!",
			})
			return
		}

		if req.Code == "" {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "google code not provided!",
			})
			return
		}

		googleToken, err := googleOAuth2Config.Exchange(ctx.Request.Context(), req.Code)
		if err != nil {
			fmt.Println("Error exchanging authorization code:", err)
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "error exchanging authorization code!",
			})
			return
		}

		client := googleOAuth2Config.Client(ctx.Request.Context(), googleToken)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			fmt.Println("Error getting user info:", err)
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "Error getting user info!",
			})
			return
		}
		content, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		var responseData googleOAuth2Response
		json.Unmarshal(content, &responseData)

		count, err := models.UserCollection.CountDocuments(ctx, bson.M{"email": responseData.Email})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Success: false,
				Message: "something isn't working!",
			})
			return
		}

		var user models.User
		var successMsg = ""
		if count > 0 {
			// login

			err := models.UserCollection.FindOne(ctx, bson.M{"email": responseData.Email}).Decode(&user)
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

			successMsg = "user logged in successfully"
		} else {
			// Signup

			user.ID = primitive.NewObjectID()
			user.Email = responseData.Email
			user.FirstName = responseData.GivenName
			user.LastName = responseData.FamilyName
			user.ProfilePicUrl = responseData.Picture
			user.IsBlocked = false
			user.Role = constants.UserRole
			user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			_, insertErr := models.UserCollection.InsertOne(ctx, user)
			if insertErr != nil {
				ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
					Success: false,
					Message: "user is not created",
				})
				return
			}

			successMsg = "user created successfully"
		}

		token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.Role)

		ctx.JSON(http.StatusOK, responses.AuthResponse{
			Success:       true,
			Message:       successMsg,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Email:         user.Email,
			Phone:         user.Phone,
			Role:          user.Role,
			IsBlocked:     user.IsBlocked,
			ProfilePicUrl: user.ProfilePicUrl,
			Token:         token,
			RefreshToken:  refreshToken,
		})
	}
}
