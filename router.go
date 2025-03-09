package dummyservice

import "net/http"

func CreateRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("OPTION /{route...}", Success)
	r.HandleFunc("OPTION /failure", BadRequest)
	r.HandleFunc("OPTION /error", ReturnErrorCode)

	r.HandleFunc("GET /success", Success)
	r.HandleFunc("GET /notfound", NotFound)
	r.HandleFunc("GET /error", BadRequest)
	r.HandleFunc("GET /error/{code}", ReturnErrorCode)
	r.HandleFunc("GET /timeout", TimeoutWithDuration)
	r.HandleFunc("GET /timeout/{duration}", TimeoutWithDuration)

	r.HandleFunc("POST /success", Success)
	r.HandleFunc("POST /notfound", NotFound)
	r.HandleFunc("POST /error", BadRequest)
	r.HandleFunc("POST /error/{code}", ReturnErrorCode)
	r.HandleFunc("POST /timeout", TimeoutWithDuration)
	r.HandleFunc("POST /timeout/{duration}", TimeoutWithDuration)

	r.HandleFunc("PUT /success", Success)
	r.HandleFunc("PUT /notfound", NotFound)
	r.HandleFunc("PUT /error", BadRequest)
	r.HandleFunc("PUT /error/{code}", ReturnErrorCode)
	r.HandleFunc("PUT /timeout", TimeoutWithDuration)
	r.HandleFunc("PUT /timeout/{duration}", TimeoutWithDuration)

	r.HandleFunc("PATCH /success", Success)
	r.HandleFunc("PATCH /notfound", NotFound)
	r.HandleFunc("PATCH /error", BadRequest)
	r.HandleFunc("PATCH /error/{code}", ReturnErrorCode)
	r.HandleFunc("PATCH /timeout", TimeoutWithDuration)
	r.HandleFunc("PATCH /timeout/{duration}", TimeoutWithDuration)

	r.HandleFunc("DELETE /success", Success)
	r.HandleFunc("DELETE /notfound", NotFound)
	r.HandleFunc("DELETE /error", BadRequest)
	r.HandleFunc("DELETE /error/{code}", ReturnErrorCode)
	r.HandleFunc("DELETE /timeout", TimeoutWithDuration)
	r.HandleFunc("DELETE /timeout/{duration}", TimeoutWithDuration)
	return r
}
