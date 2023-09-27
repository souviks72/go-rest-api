//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"testing"

	resty "github.com/go-resty/resty/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString()
}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+createToken()).
			SetBody(`{"slug":"/", "author":"Elliot", "body": "hellow world"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("cannot post comment without JWT header", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug":"/", "author":"Elliot", "body": "hellow world"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

}
