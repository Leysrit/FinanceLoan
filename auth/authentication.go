package auth

import (
	"Finance/payload"
	"Finance/utility"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

const (
	COST       = 10
	SECRET_KEY = "secret"
)

func ExtractJwtFromHeader(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}

func GetClaimsFromJwt(tokenString string) (*payload.Claims, error) {
	var claims *payload.Claims
	token, err := jwt.ParseWithClaims(tokenString, &payload.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*payload.Claims); ok && token.Valid {
			return claims, nil
		}
	}

	if !token.Valid {
		return nil, utility.ErrUnauthorized
	}

	fmt.Println("test")
	return claims, nil
}
