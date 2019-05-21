package helpers

import (
	"errors"
	"github.com/labstack/gommon/log"
	"gopkg.in/h2non/gentleman.v2"
)

type tokenResponse struct {
	// API version
	Version string       `json:"v"`
	Token   Token `json:"token"`
}

// Token access_token information
type Token struct {
	UserID      int      `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// TokenVerify check token and return info about token (role, user_id and etc)
func TokenVerify(accessToken string, required bool, roles []string, scopes []string) (Token, error) {
	req := gentleman.New().URL(
		GetEnv("SERVICE_AUTH", "http://service-auth:8080") +
		"/token.check?access_token=" + accessToken,
	).Request()
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		log.Error(err)
		return Token{}, err
	}
	if !res.Ok {
		// auth is not required
		if !required {
			return Token{}, nil
		}
		return Token{}, errors.New("access_token not founded")
	}

	var tokenres tokenResponse
	if err := res.JSON(tokenres); err != nil {
		log.Error(err)
		return Token{}, err
	}

	token := tokenres.Token

	if !ContainsString(roles, token.Role) {
		return Token{}, errors.New("user does not have the necessary roles")
	}

	if !ContainsStrings(scopes, token.Permissions) {
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
