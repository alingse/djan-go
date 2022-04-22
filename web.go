package djan

import "net/http"

func (d *djanAdmin) Handler(prefix string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
