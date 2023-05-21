package users

import (
	"time"

	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iss": "localhost:3001/login",
		"user": gin.H{
			"fName": user.FirstName,
			"lName": user.LastName,
			"email": user.Email,
			"uname": user.Username,
			"id":    user.ID,
		},
	})

	return token.SignedString(common.GetJWTPrivateKey())
}
