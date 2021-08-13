package model

import (
	db_model "gouser_center/dal/db/model"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID uint
	Extra  db_model.UserExtra
	jwt.StandardClaims
}
