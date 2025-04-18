package dummyservice

import "net/http"

func CreateRouter() *http.ServeMux {
	r := http.NewServeMux()

	// Option routes with different pattern
	r.HandleFunc("OPTION /{route...}", Success)
	r.HandleFunc("OPTION /failure", BadRequest)
	r.HandleFunc("OPTION /error", ReturnErrorCode)

	// Define routes as data structures
	type Route struct {
		pattern string
		handler http.HandlerFunc
	}

	routes := []Route{
		{"/success", Success},
		{"/notfound", NotFound},
		{"/error", BadRequest},
		{"/error/{code}", ReturnErrorCode},
		{"/timeout", TimeoutWithDuration},
		{"/timeout/{duration}", TimeoutWithDuration},
	}

	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	// Register all routes for all methods
	for _, route := range routes {
		for _, method := range methods {
			r.HandleFunc(method+" "+route.pattern, route.handler)
		}
	}

	return r
}
