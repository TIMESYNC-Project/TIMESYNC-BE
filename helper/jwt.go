package helper

import (
	"log"
	"time"
	"timesync-be/config"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 3 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, err := token.SignedString([]byte(config.JWTKey))
	if err != nil {
		log.Println(err.Error())
	}
	// log.Println(useToken, "/n", token)
	return useToken, token
}

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	userId := -1
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userId = int(claims["userID"].(float64))
		case int:
			userId = claims["userID"].(int)
		}
	}
	return userId
}
