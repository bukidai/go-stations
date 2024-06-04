package middleware

import (
	"errors"
	"net/http"
)

type AuthSetting struct {
	Username string
	Password string
}

func NewAuthSetting(username, password string) (*AuthSetting, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password must be set")

	}
	return &AuthSetting{
		Username: username,
		Password: password,
	}, nil
}

func BasicAuth(h http.Handler, authSetting *AuthSetting) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		} else if username != authSetting.Username || password != authSetting.Password {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
