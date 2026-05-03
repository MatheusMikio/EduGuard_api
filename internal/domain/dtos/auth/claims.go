package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"userId"`
	Email    string `json:"json"`
	Role     string `json:"role"`
	SchoolID uint   `json:"schoolId"`
	jwt.RegisteredClaims
}
