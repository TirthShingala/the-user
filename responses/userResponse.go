package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	IsBlocked    bool   `json:"isBlocked"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsBlocked bool   `json:"isBlocked"`
}

type AdminUserResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Role      string             `json:"role"`
	IsBlocked bool               `json:"isBlocked"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type PaginationMetaData struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"perPage"`
	PageCount  int64 `json:"pageCount"`
	TotalCount int64 `json:"totalCount"`
}

type PaginationResponse struct {
	Success  bool                `json:"success"`
	Message  string              `json:"message"`
	MetaData PaginationMetaData  `json:"metaData"`
	Data     []AdminUserResponse `json:"data"`
}
