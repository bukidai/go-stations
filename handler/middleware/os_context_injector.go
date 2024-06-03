package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

// OSContextInjector UAからOSを判定し、contextにOSを追加する

type osContextKey string

const osKey = osContextKey("os")

func OSContextInjector(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		parsedUA := useragent.Parse(r.UserAgent())
		ctx := r.Context()
		os := parsedUA.OS
		ctx = setOS(ctx, os)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func setOS(ctx context.Context, os string) context.Context {
	return context.WithValue(ctx, osKey, os)
}

func GetOS(ctx context.Context) string {
	if os, ok := ctx.Value(osKey).(string); ok {
		return os
	}
	return ""
}
