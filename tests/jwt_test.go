package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	hardcodeauth "github.com/sifatulrabbi/hardcode-auth"
)

func TestGenerateLoginToken(t *testing.T) {
	email := "sifatuli.r@gmail.com"
	token, err := hardcodeauth.GenerateLoginCookieValidationJWT(email)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	fmt.Println("SUCCESS: token generated")

	tokenEmail, err := hardcodeauth.ParseLoginCookieJWT(token)
	if err != nil || tokenEmail != email {
		t.Errorf("err=%q or token mismatch tokenEmail=%q email=%q\n", err, tokenEmail, email)
		t.Fail()
		return
	}
	fmt.Println("SUCCESS: token parsed")
}

func TestSignatureValidation(t *testing.T) {
	email := "sifatuli.r@gmail.com"
	claims := hardcodeauth.LoginCookieValidationClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenWithInvalidSig, err := token.SignedString([]byte("invalid signature"))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if _, err := hardcodeauth.ParseLoginCookieJWT(tokenWithInvalidSig); err == nil {
		t.Error("Unable to catch invalid token signature")
		t.Fail()
	}
	fmt.Println("SUCCESS: able to identify invalid tokens")
}
