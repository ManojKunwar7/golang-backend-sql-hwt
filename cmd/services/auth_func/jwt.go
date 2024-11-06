package auth_func

import (
	"strconv"
	"test-project/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	Expiration := time.Duration(config.ENVS.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"ExpiresAt": time.Now().Add(Expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT() {

}
