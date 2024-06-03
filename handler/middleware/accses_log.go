package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AccessLogInfo struct {
	TimeStamp time.Time `json:"time_stamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func AccessLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timeStamp := time.Now()
		path := r.URL.Path
		os := GetOS(r.Context())
		h.ServeHTTP(w, r)
		latency := time.Since(timeStamp).Milliseconds()
		accessLogInfo := AccessLogInfo{
			TimeStamp: timeStamp,
			Latency:   latency,
			Path:      path,
			OS:        os,
		}
		if m, err := json.Marshal(accessLogInfo); err != nil {
			log.Panicln(err)
		} else {
			fmt.Println(string(m))
		}

	}
	return http.HandlerFunc(fn)
}
