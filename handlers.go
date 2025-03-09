package dummyservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Success(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(404), 404)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(400), 400)
}

func ReturnErrorCode(w http.ResponseWriter, r *http.Request) {
	codeString := r.PathValue("code")
	code, err := strconv.Atoi(codeString)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}

func TimeoutWithDuration(w http.ResponseWriter, r *http.Request) {
	durationString := r.PathValue("duration")
	duration, err := strconv.Atoi(durationString)
	if err != nil {
		duration = 30
	}
	time.Sleep(time.Duration(duration) * time.Second)
	w.Write([]byte(fmt.Sprintf("done after %d seconds", duration)))
}
