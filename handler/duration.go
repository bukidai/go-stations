package handler

import (
	"net/http"
	"strconv"
	"time"
)

type DurationHandler struct{}

func NewDurationHandler() *DurationHandler {
	return &DurationHandler{}
}

// /duration?duration=secで指定されたミリ秒数だけ待機するエンドポイント
func (h *DurationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { // GETメソッド以外は405 Method Not Allowed
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	query := r.URL.Query()
	duration := query.Get("duration")
	if duration == "" { // durationが空文字の場合は400 Bad Request
		http.Error(w, "duration is required", http.StatusBadRequest)
		return
	}
	if durationSec, err := strconv.Atoi(duration); err != nil {
		http.Error(w, "duration must be integer", http.StatusBadRequest)
		return
	} else {
		time.Sleep(time.Duration(durationSec) * time.Millisecond)
		msg := "waited for " + duration + " milliseconds\n"
		w.Write([]byte(msg))
	}
}
