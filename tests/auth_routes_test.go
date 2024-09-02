package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", http.NoBody)
	failOnErr(err)
	res, err := http.DefaultClient.Do(req)
	failOnErr(err)
	expectedRes := struct {
		user        interface{}
		accessToken string
	}{}
	defer res.Body.Close()
	failOnErr(json.NewDecoder(res.Body).Decode(&expectedRes))
}

func failOnErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
