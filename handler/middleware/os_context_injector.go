package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

// OSContextInjector UAからOSを判定し、contextにOSを追加する

type OSKey string

func OSContextInjector(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		parsedUA := useragent.Parse(r.UserAgent())
		ctx := r.Context()
		ctx = context.WithValue(ctx, OSKey("os"), parsedUA.OS)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
