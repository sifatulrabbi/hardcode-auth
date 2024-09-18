package hardcodeauth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func init() {
	prepareConfigs()
}

type LoginCookieValidationClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateLoginCookieValidationJWT(email string) (string, error) {
	claims := LoginCookieValidationClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenStr, err := token.SignedString([]byte(ENVConfig.JWT_SECRET)); err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func ParseLoginCookieJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		expTime, err := t.Claims.GetExpirationTime()
		if err != nil {
			return nil, err
		}
		if expTime.Before(time.Now()) {
			return nil, fmt.Errorf("The token has expired!")
		}
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(ENVConfig.JWT_SECRET), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return "", errors.New("invalid token claims")
	} else if email, ok := claims["email"].(string); !ok {
		return "", errors.New("invalid token claims")
	} else {
		return email, nil
	}
}
