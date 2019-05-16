package helpers

import (
	"errors"
	"github.com/labstack/gommon/log"
	"gopkg.in/h2non/gentleman.v2"
)

// Token access_token information
type Token struct {
	UserID      int      `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// TokenRequired check token and return info about token (role, user_id and etc)
func TokenVerify(accessToken string, required bool, roles []string, scopes []string) (Token, error) {
	req := gentleman.New().URL(GetEnv("SERVICE_AUTH", "http://service-auth:8080")).Request()
	req.Path("/token.check?access_token=" + accessToken)

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
		return Token{}, errors.New(res.RawResponse.Status)
	}

	var token Token
	if err := res.JSON(token); err != nil {
		log.Error(err)
		return Token{}, err
	}

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

// ContainsString returns true if a strings is present in a iteratee.
func ContainsStrings(s []string, v []string) bool {
	count := 0
	for _, v1 := range s {
		if ContainsString(v, v1) {
			count += 1
		}
	}
	return count == len(v)
}
