package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var tokenSecret = []byte("random-string")

func Generatetoken(uid string) (string, error) {
	var err error

	//create a map for claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	//create a token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Generate a token string
	tokenString, err := token.SignedString(tokenSecret)

	if err != nil {
		err = fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func parseJWT(header string, r *http.Request) (*jwt.Token, error) {
	var tokenStr = getToken(r)
	_, err := fmt.Scanf(header, "Bearer %s", &tokenStr)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])
		}
		return []byte("random_string"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func getToken(r *http.Request) string {
	beartoken := r.Header.Get("Authorization")

	strArr := strings.Split(beartoken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
