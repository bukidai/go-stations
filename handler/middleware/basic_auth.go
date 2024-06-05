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
		if !ok || (username != authSetting.Username || password != authSetting.Password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Auth area"`)
			http.Error(w, "401 authentication required", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
