package helpers

import (
	"errors"
	"github.com/labstack/gommon/log"
	"gopkg.in/resty.v1"
	"net/http"
)

type tokenResponse struct {
	// API version
	Version string `json:"v"`
	Token   Token  `json:"token"`
}

// Token access_token information
type Token struct {
	UserID      int      `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// TokenVerify check token and return info about token (role, user_id and etc)
func TokenVerify(accessToken string, required bool, roles []string, scopes []string) (Token, error) {
	var tokenres tokenResponse
	resp, err := resty.R().
		SetResult(&tokenres).
		Get(GetEnv("SERVICE_AUTH", "http://service-auth:8080") + "/token.check?access_token=" + accessToken)
	if err != nil {
		log.Error(err)
		return Token{}, err
	}

	if resp.StatusCode() == http.StatusNotFound || resp.StatusCode() == http.StatusUnauthorized {
		// auth is not required
		if !required {
			return Token{}, nil
		}
		return Token{}, errors.New("access_token not founded")
	}

	token := tokenres.Token
	if !ContainsString(roles, token.Role) {
		return Token{}, errors.New("user does not have the necessary roles")
	}

	if !ContainsStrings(token.Permissions, scopes) {
		return Token{}, errors.New("user does not have the necessary permissions(scopes)")
	}

	return token, nil
}

// ContainsString returns true if a string is present in a iteratee.
func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainsStrings returns true if a strings is present in a iteratee.
func ContainsStrings(s []string, v []string) bool {
	count := 0
	for _, v1 := range s {
		if ContainsString(v, v1) {
			count++
		}
	}
	return count == len(v)
}
