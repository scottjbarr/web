package web

import "net/http"

// Healthcheck provides a rudimentary "healthcheck" handler.
func Healthcheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}
