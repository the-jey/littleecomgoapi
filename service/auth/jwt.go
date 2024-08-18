package auth

import "github.com/golang-jwt/jwt/v5"

func CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{})
}
